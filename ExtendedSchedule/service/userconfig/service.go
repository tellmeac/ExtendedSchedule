package userconfig

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/repository"
)

func NewService(repository repository.IUserConfigRepository) IService {
	return &service{repository: repository}
}

type IService interface {
	GetUserConfig(ctx context.Context, email string) (aggregate.UserConfig, error)
	UpdateUserConfig(ctx context.Context, desired aggregate.UserConfig) error
}

type service struct {
	repository repository.IUserConfigRepository
}

func (s service) GetUserConfig(ctx context.Context, email string) (aggregate.UserConfig, error) {
	config, err := s.repository.GetByEmail(ctx, email)
	switch {
	case errors.Is(err, repository.ErrConfigNotFound):
		newConfig := aggregate.NewUserConfig(email)
		if err := s.repository.Put(ctx, newConfig); err != nil {
			return aggregate.UserConfig{}, fmt.Errorf("failed to create new config: %w", err)
		}
		return newConfig, nil
	case err != nil:
		return aggregate.UserConfig{}, err
	default:
		return config, nil
	}
}

func (s service) UpdateUserConfig(ctx context.Context, desired aggregate.UserConfig) error {
	return s.repository.Update(ctx, desired)
}
