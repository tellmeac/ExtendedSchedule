package userconfig

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	commonmodels "tellmeac/extended-schedule/common/models"
	errs "tellmeac/extended-schedule/lib/errors"
	"tellmeac/extended-schedule/pkg/userconfig/dao"
)

var Module = fx.Options(
	fx.Provide(dao.New),
	fx.Provide(New),
)

// Manager provides methods to work with user configuration.
type Manager interface {
	GetByEmail(ctx context.Context, email string) (*commonmodels.UserConfig, error)
	Update(ctx context.Context, desired *commonmodels.UserConfig) error
	Onboard(ctx context.Context, config *commonmodels.UserConfig) error
}

// New returns a default implementation of Manager.
func New(dao dao.DAO) Manager {
	return &manager{dao: dao}
}

type manager struct {
	dao dao.DAO
}

func (m manager) GetByEmail(ctx context.Context, email string) (*commonmodels.UserConfig, error) {
	config, err := m.dao.GetByEmail(ctx, email)

	switch {
	case errors.Is(err, errs.ErrNotFound):
		log.Info().Str("email", email).Msg("User config was not found, will be created new")
		newConfig := commonmodels.NewUserConfig(email)
		err = m.Onboard(ctx, &newConfig)
		return &newConfig, err
	case err != nil:
		return nil, err
	default:
		return config, nil
	}
}

func (m manager) Update(ctx context.Context, desired *commonmodels.UserConfig) error {
	return m.dao.Update(ctx, desired)
}

func (m manager) Onboard(ctx context.Context, config *commonmodels.UserConfig) error {
	return m.dao.Onboard(ctx, config)
}
