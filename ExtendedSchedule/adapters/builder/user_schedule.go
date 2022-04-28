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
		schedule: schedule,
		configs:  configs,
	}
}

type UserScheduleBuilder struct {
	schedule providers.IBaseScheduleProvider
	configs  repository.IUserConfigRepository
}

func (builder UserScheduleBuilder) Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	//TODO implement me
	panic("implement me")
}
