package userconfig

import (
	"context"
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/aggregates"
	"tellmeac/extended-schedule/domain/repository"
)

func NewService(repository repository.IUserConfigRepository) IService {
	return &Service{repository: repository}
}

type IService interface {
	GetUserConfig(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error)
	UpdateUserConfig(ctx context.Context, desired aggregates.UserConfig) error
}

type Service struct {
	repository repository.IUserConfigRepository
}

func (s Service) GetUserConfig(ctx context.Context, userID uuid.UUID) (aggregates.UserConfig, error) {
	return s.repository.Get(ctx, userID)
}

func (s Service) UpdateUserConfig(ctx context.Context, desired aggregates.UserConfig) error {
	return s.repository.Update(ctx, desired.UserID, desired)
}
