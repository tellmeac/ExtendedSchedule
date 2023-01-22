package adapters

import (
	"context"
	"fmt"
	"tellmeac/extended-schedule/adapters/ent"
	uc "tellmeac/extended-schedule/adapters/ent/userconfig"
	"tellmeac/extended-schedule/pkg/errors"
	"tellmeac/extended-schedule/userconfig"
)

func NewUserConfigRepository(client *ent.Client) UserConfigRepository {
	return UserConfigRepository{client: client}
}

type UserConfigRepository struct {
	client *ent.Client
}

func (r UserConfigRepository) GetByEmail(ctx context.Context, email string) (userconfig.UserConfig, error) {
	dbo, err := r.client.UserConfig.Query().Where(uc.Email(email)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return userconfig.UserConfig{}, fmt.Errorf("user config was not found by email: %w", errors.ErrNotFound)
	case err != nil:
		return userconfig.UserConfig{}, err
	}

	return userconfig.UserConfig{
		ID:             dbo.ID,
		Email:          dbo.Email,
		Base:           *dbo.Base,
		ExtendedGroups: dbo.ExtendedGroups,
		ExcludeRules:   dbo.ExcludeRules,
	}, nil
}

func (r UserConfigRepository) Upsert(ctx context.Context, email string, desired userconfig.UserConfig) error {
	return r.client.UserConfig.Create().
		SetEmail(email).
		SetBase(&desired.Base).
		SetExtendedGroups(desired.ExtendedGroups).
		SetExcludeRules(desired.ExcludeRules).
		OnConflict().UpdateNewValues().Exec(ctx)
}
