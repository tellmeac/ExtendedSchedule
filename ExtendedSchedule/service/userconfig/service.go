package userconfig

import (
	"context"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/aggregates"
	"github.com/tellmeac/extended-schedule/domain/entity"
)

type IService interface {
	GetUserConfig(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
	UpdateJoinedGroups(ctx context.Context, userID uuid.UUID, desired entity.JoinedGroups) error
	UpdateIgnoredLessons(ctx context.Context, userID uuid.UUID, desired []entity.ExcludedLesson) error
	UpdateExtendedLessons(ctx context.Context, userID uuid.UUID, desired []entity.ExtendedLesson) error
}
