package repository

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/ent"
	"tellmeac/extended-schedule/adapters/ent/userconfig"
	"tellmeac/extended-schedule/domain/aggregate"
	"tellmeac/extended-schedule/domain/repository"
)

// NewEntUserConfigRepository создает репозиторий для пользовательской конфигурации.
func NewEntUserConfigRepository(client *ent.Client) repository.IUserConfigRepository {
	return &entUserConfigRepository{
		client: client,
	}
}

type entUserConfigRepository struct {
	client *ent.Client
}

func (r entUserConfigRepository) Put(ctx context.Context, userConfig aggregate.UserConfig) error {
	_, err := r.client.UserConfig.Create().
		SetEmail(userConfig.Email).
		SetBaseGroup(userConfig.BaseGroup).
		SetExtendedGroupLessons(userConfig.ExtendedGroupLessons).
		SetExcludedLessons(userConfig.ExcludedLessons).
		Save(ctx)
	return err
}

func (r entUserConfigRepository) GetByEmail(ctx context.Context, email string) (aggregate.UserConfig, error) {
	dbo, err := r.client.UserConfig.Query().Where(userconfig.EmailEqualFold(email)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return aggregate.UserConfig{}, fmt.Errorf("user = %s: %w", email, repository.ErrConfigNotFound)
	case err != nil:
		return aggregate.UserConfig{}, err
	default:
		return aggregate.UserConfig{
			Email:                dbo.Email,
			BaseGroup:            dbo.BaseGroup,
			ExcludedLessons:      dbo.ExcludedLessons,
			ExtendedGroupLessons: dbo.ExtendedGroupLessons,
		}, nil
	}
}

func (r entUserConfigRepository) Update(ctx context.Context, desired aggregate.UserConfig) error {
	_, err := r.client.UserConfig.Update().
		Where(userconfig.EmailEqualFold(desired.Email)).
		SetBaseGroup(desired.BaseGroup).
		SetExtendedGroupLessons(desired.ExtendedGroupLessons).
		SetExcludedLessons(desired.ExcludedLessons).
		Save(ctx)
	return err
}
