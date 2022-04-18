package factory

import (
	"context"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/aggregates"
)

type IUserConfigFactory interface {
	Make(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
}
