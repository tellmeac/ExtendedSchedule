package userconfig

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type IService interface {
	GetUserConfig(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
	UpdateJoinedGroups(ctx context.Context, userID uuid.UUID, desired entity.JoinedGroups) error
	UpdateIgnoredLessons(ctx context.Context, userID uuid.UUID, desired []entity.ExcludedLesson) error
	UpdateExtendedLessons(ctx context.Context, userID uuid.UUID, desired []entity.ExtendedLesson) error
}
