package providers

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"
	"time"
)

type IBaseScheduleProvider interface {
	GetByGroup(ctx context.Context, groupID string, start time.Time, end time.Time) ([]entity.DaySchedule, error)
}
