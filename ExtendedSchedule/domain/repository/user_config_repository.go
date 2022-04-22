package repository

import (
	"context"
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type IJoinedGroupsRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.GroupInfo, error)
}
