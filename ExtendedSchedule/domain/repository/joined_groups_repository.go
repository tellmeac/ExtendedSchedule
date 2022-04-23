package repository

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type IJoinedGroupsRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.GroupInfo, error)
	Update(ctx context.Context, userID uuid.UUID, desired []entity.GroupInfo) error
}
