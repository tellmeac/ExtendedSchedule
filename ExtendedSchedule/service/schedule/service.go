package schedule

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/builder"
	"tellmeac/extended-schedule/domain/provider"
	"time"
)

type IService interface {
	GetPersonal(ctx context.Context, userIdentifier string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
	GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
	GetByLesson(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
}

// NewService создает сервис для получения расписания.
func NewService(schedule provider.IBaseScheduleProvider, builder builder.IUserScheduleBuilder) IService {
	return &Service{
		schedule: schedule,
		builder:  builder,
	}
}

type Service struct {
	schedule provider.IBaseScheduleProvider
	builder  builder.IUserScheduleBuilder
}

func (s Service) GetPersonal(ctx context.Context, userIdentifier string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error) {
	return s.builder.Make(ctx, userIdentifier, start, end)
}

func (s Service) GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error) {
	return s.schedule.GetByGroupID(ctx, groupID, start, end)
}

func (s Service) GetByLesson(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error) {
	return s.schedule.GetByLessonID(ctx, groupID, lessonID, start, end)
}
