package repository

import (
	"context"
	"errors"
	"tellmeac/extended-schedule/domain/aggregate"
)

var ErrConfigNotFound = errors.New("user config was not found")

// IUserConfigRepository предоставляет методы для работы с aggregates.UserConfig.
type IUserConfigRepository interface {
	Put(ctx context.Context, config aggregate.UserConfig) error
	GetByEmail(ctx context.Context, userEmail string) (aggregate.UserConfig, error)
	Update(ctx context.Context, desired aggregate.UserConfig) error
}
