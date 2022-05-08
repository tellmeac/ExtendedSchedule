package repository

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregate"
)

// IUserConfigRepository предоставляет методы для работы с aggregates.UserConfig.
type IUserConfigRepository interface {
	Get(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error)
	Update(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error
}
