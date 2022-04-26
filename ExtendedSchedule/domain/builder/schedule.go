package builder

import (
	"context"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/aggregates"
	"time"
)

// IUserScheduleBuilder строит польозвательское расписание.
type IUserScheduleBuilder interface {
	// Make формирует расписание пользователя.
	Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]aggregates.DaySchedule, error)
}
