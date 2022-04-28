package builder

import (
	"context"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/builder"
	"tellmeac/extended-schedule/domain/providers"
	"tellmeac/extended-schedule/domain/repository"
	"time"
)

func NewUserScheduleBuilder(schedule providers.IBaseScheduleProvider, configs repository.IUserConfigRepository) builder.IUserScheduleBuilder {
	return &UserScheduleBuilder{
		scheduleProvider: schedule,
		configs:          configs,
	}
}

type UserScheduleBuilder struct {
	scheduleProvider providers.IBaseScheduleProvider
	configs          repository.IUserConfigRepository
}

func (builder UserScheduleBuilder) Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	var schedule []aggregates.DaySchedule = nil

	config, err := builder.configs.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	for _, group := range config.JoinedGroups {
		groupSchedule, err := builder.scheduleProvider.GetByGroupID(ctx, group.ID, start, end)
		if err != nil {
			return nil, err
		}

		schedule, err = aggregates.JoinSchedules(schedule, groupSchedule)
		if err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(schedule); i++ {
		if err := schedule[i].ExcludeLessons(config.ExcludedLessons); err != nil {
			return nil, err
		}
	}

	return schedule, nil
}
