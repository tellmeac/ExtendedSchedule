package userconfig

import (
	"context"
	"github.com/pkg/errors"
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
	config, err := s.repository.Get(ctx, userIdentifier)
	switch {
	case errors.Is(err, repository.ErrConfigNotFound):
		return s.repository.Init(ctx, userIdentifier)
	case err != nil:
		return aggregate.UserConfig{}, err
	default:
		return config, nil
	}
}

func (s Service) UpdateUserConfig(ctx context.Context, userIdentifier string, desired aggregate.UserConfig) error {
	return s.repository.Update(ctx, userIdentifier, desired)
}
