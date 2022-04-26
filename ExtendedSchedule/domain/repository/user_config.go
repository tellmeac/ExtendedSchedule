package repository

import (
	"context"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/aggregates"
)

// IUserConfigRepository предоставляет методы для работы с aggregates.UserConfig.
type IUserConfigRepository interface {
	Get(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
	Update(ctx context.Context, userID uuid.UUID, desired aggregates.UserConfig) error
}
