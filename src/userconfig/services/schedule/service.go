package schedule

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"github.com/tellmeac/extended-schedule/userconfig/domain/schedule"
	"github.com/tellmeac/extended-schedule/userconfig/domain/userconfig"
	"github.com/tellmeac/extended-schedule/userconfig/services/tsuschedule"
	configservice "github.com/tellmeac/extended-schedule/userconfig/services/userconfig"
	"time"
)

const scheduleDateFormat = "2006-01-02"

// Service interface with methods to get schedule.
type Service interface {
	GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]schedule.DaySchedule, error)
	GetUserScheduleByEmail(ctx context.Context, email string, start time.Time, end time.Time) ([]schedule.DaySchedule, error)
}

// New returns default Service implementation.
func New(client *tsuschedule.Client, configs configservice.Service) Service {
	return &service{
		client:  client,
		configs: configs,
	}
}

type service struct {
	client  *tsuschedule.Client
	configs configservice.Service
}

func (s service) GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]schedule.DaySchedule, error) {
	params := tsuschedule.GetScheduleGroupParams{
		Id:       groupID,
		DateFrom: start.Format(scheduleDateFormat),
		DateTo:   end.Format(scheduleDateFormat),
	}

	response, err := s.client.GetScheduleGroup(ctx, &params)
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

	var result = make([]schedule.DaySchedule, len(*scheduleDto.JSON200))
	for i, day := range *scheduleDto.JSON200 {
		result[i], err = toCommonDay(day)
		if err != nil {
			return nil, fmt.Errorf("failed to map response data properly: %w", err)
		}
	}

	return result, nil
}

func (s service) GetUserScheduleByEmail(ctx context.Context, email string, start time.Time, end time.Time) ([]schedule.DaySchedule, error) {
	var userSchedule []schedule.DaySchedule = nil

	config, err := s.configs.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	var baseSchedule = make([]schedule.DaySchedule, 0)
	if config.BaseGroup != nil {
		baseSchedule, err = s.GetByGroup(ctx, config.BaseGroup.ExternalID, start, end)
		if err != nil {
			return nil, err
		}
	}

	userSchedule, err = schedule.Join(userSchedule, baseSchedule)
	if err != nil {
		return nil, err
	}

	extended, err := s.getExtendedSchedule(ctx, config.ExtendedGroupLessons, start, end)
	if err != nil {
		return nil, err
	}

	userSchedule, err = schedule.Join(userSchedule, extended)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(userSchedule); i++ {
		if err := userSchedule[i].ExcludeLessons(config.ExcludeRules); err != nil {
			return nil, err
		}
	}

	return userSchedule, nil
}

func (s service) getExtendedSchedule(ctx context.Context, groupLessons []userconfig.ExtendedGroupLessons, start time.Time, end time.Time) ([]schedule.DaySchedule, error) {
	var result []schedule.DaySchedule
	for _, extended := range groupLessons {
		log.Info().Str("group", extended.Group.ExternalID).Msg("Apply extended group lessons to user schedule")
		groupSchedule, err := s.GetByGroup(ctx, extended.Group.ExternalID, start, end)
		if err != nil {
			return nil, err
		}

		filtered := filterByLessons(groupSchedule, extended.LessonIDs)
		result, err = schedule.Join(result, filtered)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

// filterByLessons leaves only the necessary items to the schedule.
func filterByLessons(s []schedule.DaySchedule, lessonIDs []string) []schedule.DaySchedule {
	var shouldBeIncluded = make(map[string]bool)
	for _, id := range lessonIDs {
		shouldBeIncluded[id] = true
	}

	var result = make([]schedule.DaySchedule, 0, len(s))
	for _, day := range s {
		result = append(result, schedule.DaySchedule{
			Date: day.Date,
			Lessons: lo.Filter(day.Lessons, func(lesson schedule.Lesson, _ int) bool {
				if _, ok := shouldBeIncluded[lesson.ID]; ok {
					return true
				}
				return false
			}),
		})
	}
	return result
}
