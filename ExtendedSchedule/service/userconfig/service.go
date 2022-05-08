package userconfig

import (
	"context"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/repository"
)

func NewService(repository repository.IUserConfigRepository) IService {
	return &Service{repository: repository}
}

type IService interface {
	GetUserConfig(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error)
	UpdateUserConfig(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error
}

type Service struct {
	repository repository.IUserConfigRepository
}

func (s Service) GetUserConfig(ctx context.Context, userIdentifier string) (aggregate.UserConfig, error) {
	return s.repository.Get(ctx, userIdentifier)
}

func (s Service) UpdateUserConfig(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error {
	return s.repository.Update(ctx, userIdentifier, desired)
}
