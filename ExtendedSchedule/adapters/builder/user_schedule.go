package builder

import (
	"context"
	"github.com/samber/lo"
	"tellmeac/extended-schedule/adapters/provider"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/builder"
	"tellmeac/extended-schedule/domain/entity"
	"tellmeac/extended-schedule/domain/repository"
	"time"
)

func NewUserScheduleBuilder(schedule provider.IBaseScheduleProvider, configs repository.IUserConfigRepository) builder.IUserScheduleBuilder {
	return &UserScheduleBuilder{
		scheduleProvider: schedule,
		configs:          configs,
	}
}

type UserScheduleBuilder struct {
	scheduleProvider provider.IBaseScheduleProvider
	configs          repository.IUserConfigRepository
}

func (builder UserScheduleBuilder) Make(ctx context.Context, userEmail string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error) {
	var schedule []aggregate.DaySchedule = nil

	config, err := builder.configs.GetByEmail(ctx, userEmail)
	if err != nil {
		return nil, err
	}

	var baseSchedule = make([]aggregate.DaySchedule, 0)
	if config.BaseGroup != nil {
		baseSchedule, err = builder.scheduleProvider.GetByGroupID(ctx, config.BaseGroup.ID, start, end)
		if err != nil {
			return nil, err
		}
	}

	schedule, err = aggregate.JoinSchedules(schedule, baseSchedule)
	if err != nil {
		return nil, err
	}

	extended, err := builder.getExtendedSchedule(ctx, config.ExtendedGroupLessons, start, end)
	if err != nil {
		return nil, err
	}

	schedule, err = aggregate.JoinSchedules(schedule, extended)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(schedule); i++ {
		if err := schedule[i].ExcludeLessons(config.ExcludedLessons); err != nil {
			return nil, err
		}
	}

	return schedule, nil
}

func (builder UserScheduleBuilder) getExtendedSchedule(ctx context.Context, groupLessons []aggregate.ExtendedGroupLessons, start time.Time, end time.Time) ([]aggregate.DaySchedule, error) {
	var result []aggregate.DaySchedule
	for _, extended := range groupLessons {
		groupSchedule, err := builder.scheduleProvider.GetByGroupID(ctx, extended.Group.ID, start, end)
		if err != nil {
			return nil, err
		}

		filtered := filterByLessons(groupSchedule, extended.LessonIDs)
		result, err = aggregate.JoinSchedules(result, filtered)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func filterByLessons(schedule []aggregate.DaySchedule, lessonIDs []string) []aggregate.DaySchedule {
	var shouldBeIncluded = make(map[string]bool)
	for _, id := range lessonIDs {
		shouldBeIncluded[id] = true
	}

	var result = make([]aggregate.DaySchedule, 0, len(schedule))
	for _, day := range schedule {
		result = append(result, aggregate.DaySchedule{
			Date: day.Date,
			Lessons: lo.Filter(day.Lessons, func(lesson entity.Lesson, _ int) bool {
				if _, ok := shouldBeIncluded[lesson.ID]; ok {
					return true
				}
				return false
			}),
		})
	}
	return result
}
