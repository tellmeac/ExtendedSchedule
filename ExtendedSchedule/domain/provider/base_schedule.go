package provider

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregate"
	"time"
)

// IBaseScheduleProvider представляет провайдер для получения общего расписания.
type IBaseScheduleProvider interface {
	GetByGroupID(ctx context.Context, groupID string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
	GetByLessonID(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
}
