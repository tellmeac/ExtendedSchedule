package providers

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type IBaseScheduleProvider interface {
	GetByGroupID(ctx context.Context, groupID string, start time.Time, end time.Time) ([]entity.DaySchedule, error)
	GetLessonSchedule(ctx context.Context, groupID string, lessonID string, start time.Time, end time.Time) ([]entity.DaySchedule, error)
}
