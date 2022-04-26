package providers

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregates"
	"time"
)

// IBaseScheduleProvider представляет провайдер для получения общего расписания.
type IBaseScheduleProvider interface {
	GetByGroupID(ctx context.Context, groupID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
	GetLessonSchedule(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
}
