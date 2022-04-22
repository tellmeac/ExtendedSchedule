package schedule

import (
	"context"
	"github.com/google/uuid"
)

type IService interface {
	GetPersonal(ctx context.Context, userID uuid.UUID)
	GetByGroup(ctx context.Context, groupID uuid.UUID)
}
