package repository

import (
	"context"
	"errors"
	"tellmeac/extended-schedule/domain/aggregate"
)

var ErrConfigNotFound = errors.New("user config was not found")

// IUserConfigRepository предоставляет методы для работы с aggregates.UserConfig.
type IUserConfigRepository interface {
	Init(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error)
	Get(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error)
	Update(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error
}
