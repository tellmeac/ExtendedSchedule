package factory

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregates"

	"github.com/google/uuid"
)

type IUserConfigFactory interface {
	Make(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
}
