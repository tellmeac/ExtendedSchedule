package builder

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregate"
	"time"
)

// IUserScheduleBuilder строит польозвательское расписание.
type IUserScheduleBuilder interface {
	// Make формирует расписание пользователя.
	Make(ctx context.Context, userIdentifier string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
}
