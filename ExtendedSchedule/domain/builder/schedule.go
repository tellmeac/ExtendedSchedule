package builder

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregate"
	"time"
)

// IUserScheduleBuilder предоставляет методы для составления пользовательского расписания.
type IUserScheduleBuilder interface {
	// Make формирует расписание пользователя.
	Make(ctx context.Context, userIdentifier string, start time.Time, end time.Time) ([]aggregate.DaySchedule, error)
}
