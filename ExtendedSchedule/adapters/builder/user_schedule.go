package builder

import (
	"context"
	"tellmeac/extended-schedule/adapters/provider"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/builder"
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

	// TODO: join schedule of extended lessons

	for i := 0; i < len(schedule); i++ {
		if err := schedule[i].ExcludeLessons(config.ExcludedLessons); err != nil {
			return nil, err
		}
	}

	return schedule, nil
}
