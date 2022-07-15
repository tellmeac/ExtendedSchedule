package userconfig

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/tellmeac/extended-schedule/userconfig/dao"
	"github.com/tellmeac/extended-schedule/userconfig/domain/userconfig"
	errs "github.com/tellmeac/extended-schedule/userconfig/middle/errors"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(dao.NewUserConfigDAO),
	fx.Provide(New),
)

// Service provides methods to work with user configuration.
type Service interface {
	GetByEmail(ctx context.Context, email string) (*userconfig.UserConfig, error)
	Update(ctx context.Context, desired *userconfig.UserConfig) error
	Onboard(ctx context.Context, config *userconfig.UserConfig) error
}

// New returns a default implementation of Service.
func New(dao dao.UserConfigDAO) Service {
	return &service{dao: dao}
}

type service struct {
	dao dao.UserConfigDAO
}

func (s service) GetByEmail(ctx context.Context, email string) (*userconfig.UserConfig, error) {
	config, err := s.dao.GetByEmail(ctx, email)

	switch {
	case errors.Is(err, errs.ErrNotFound):
		log.Info().Str("email", email).Msg("User config was not found, will be created new")

		newConfig := userconfig.NewUserConfig(email)
		err = s.Onboard(ctx, &newConfig)

		return &newConfig, err
	case err != nil:
		return nil, err
	default:
		return config, nil
	}
}

func (s service) Update(ctx context.Context, desired *userconfig.UserConfig) error {
	return s.dao.Update(ctx, desired)
}

func (s service) Onboard(ctx context.Context, config *userconfig.UserConfig) error {
	return s.dao.Onboard(ctx, config)
}
