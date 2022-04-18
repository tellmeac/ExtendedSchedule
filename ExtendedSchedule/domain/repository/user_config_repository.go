package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/entity"
)

type IJoinedGroupsRepository interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]entity.GroupInfo, error)
}
