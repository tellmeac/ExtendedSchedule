package schedule

import (
	"context"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/builder"
	"tellmeac/extended-schedule/domain/providers"
	"time"
)

type IService interface {
	GetPersonal(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
	GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
	GetByLesson(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
}

// NewService создает сервис для получения расписания.
func NewService(schedule providers.IBaseScheduleProvider, builder builder.IUserScheduleBuilder) IService {
	return &Service{
		schedule: schedule,
		builder:  builder,
	}
}

type Service struct {
	schedule providers.IBaseScheduleProvider
	builder  builder.IUserScheduleBuilder
}

func (s Service) GetPersonal(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	return s.builder.Make(ctx, userID, start, end)
}

func (s Service) GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	return s.schedule.GetByGroupID(ctx, groupID, start, end)
}

func (s Service) GetByLesson(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error) {
	return s.schedule.GetByLessonID(ctx, groupID, lessonID, start, end)
}
