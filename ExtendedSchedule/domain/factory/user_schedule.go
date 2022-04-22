package factory

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"
	"time"

	"github.com/google/uuid"
)

type IUserScheduleFactory interface {
	Make(ctx context.Context, userID uuid.UUID, start time.Time, end time.Time) ([]entity.DaySchedule, error)
}
