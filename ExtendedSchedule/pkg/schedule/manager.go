package schedule

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"go.uber.org/fx"
	commonmodels "tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/pkg/tsuschedule"
	"tellmeac/extended-schedule/pkg/userconfig"
	"time"
)

var Module = fx.Options(fx.Provide(New))

const scheduleDateFormat = "2006-01-02"

// Manager interface with methods to get schedule.
type Manager interface {
	GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]commonmodels.DaySchedule, error)
	GetUserScheduleByEmail(ctx context.Context, email string, start time.Time, end time.Time) ([]commonmodels.DaySchedule, error)
}

// New returns default Manager implementation.
func New(client *tsuschedule.Client, configs userconfig.Manager) Manager {
	return &manager{
		client:  client,
		configs: configs,
	}
}

type manager struct {
	client  *tsuschedule.Client
	configs userconfig.Manager
}

func (m manager) GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]commonmodels.DaySchedule, error) {
	params := tsuschedule.GetScheduleGroupParams{
		Id:       groupID,
		DateFrom: start.Format(scheduleDateFormat),
		DateTo:   end.Format(scheduleDateFormat),
	}

	response, err := m.client.GetScheduleGroup(ctx, &params)
	if err != nil {
		return nil, fmt.Errorf("failed to get response from api: %w", err)
	}

	scheduleDto, err := tsuschedule.ParseGetScheduleGroupResponse(response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response from api: %w", err)
	}

	if scheduleDto.JSON200 == nil {
		return nil, fmt.Errorf("failed to get schedule from parsed response")
	}

	var result = make([]commonmodels.DaySchedule, len(*scheduleDto.JSON200))
	for i, day := range *scheduleDto.JSON200 {
		result[i], err = toCommonDay(day)
		if err != nil {
			return nil, fmt.Errorf("failed to map response data properly: %w", err)
		}
	}

	return result, nil
}

func (m manager) GetUserScheduleByEmail(ctx context.Context, email string, start time.Time, end time.Time) ([]commonmodels.DaySchedule, error) {
	var schedule []commonmodels.DaySchedule = nil

	config, err := m.configs.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	var baseSchedule = make([]commonmodels.DaySchedule, 0)
	if config.BaseGroup != nil {
		baseSchedule, err = m.GetByGroup(ctx, config.BaseGroup.ID, start, end)
		if err != nil {
			return nil, err
		}
	}

	schedule, err = commonmodels.JoinSchedules(schedule, baseSchedule)
	if err != nil {
		return nil, err
	}

	extended, err := m.getExtendedSchedule(ctx, config.ExtendedGroupLessons, start, end)
	if err != nil {
		return nil, err
	}

	schedule, err = commonmodels.JoinSchedules(schedule, extended)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(schedule); i++ {
		if err := schedule[i].ExcludeLessons(config.ExcludeRules); err != nil {
			return nil, err
		}
	}

	return schedule, nil
}

func (m manager) getExtendedSchedule(ctx context.Context, groupLessons []commonmodels.ExtendedGroupLessons, start time.Time, end time.Time) ([]commonmodels.DaySchedule, error) {
	var result []commonmodels.DaySchedule
	for _, extended := range groupLessons {
		log.Info().Str("group", extended.Group.ID).Msg("Apply extended group lessons to user schedule")
		groupSchedule, err := m.GetByGroup(ctx, extended.Group.ID, start, end)
		if err != nil {
			return nil, err
		}

		filtered := filterByLessons(groupSchedule, extended.LessonIDs)
		result, err = commonmodels.JoinSchedules(result, filtered)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// filterByLessons leaves only the necessary items to the schedule.
func filterByLessons(schedule []commonmodels.DaySchedule, lessonIDs []string) []commonmodels.DaySchedule {
	var shouldBeIncluded = make(map[string]bool)
	for _, id := range lessonIDs {
		shouldBeIncluded[id] = true
	}

	var result = make([]commonmodels.DaySchedule, 0, len(schedule))
	for _, day := range schedule {
		result = append(result, commonmodels.DaySchedule{
			Date: day.Date,
			Lessons: lo.Filter(day.Lessons, func(lesson commonmodels.LessonWithContext, _ int) bool {
				if _, ok := shouldBeIncluded[lesson.ID]; ok {
					return true
				}
				return false
			}),
		})
	}
	return result
}
