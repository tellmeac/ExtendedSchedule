package dao

import (
	"context"
	"fmt"
	"github.com/tellmeac/ExtendedSchedule/userconfig/domain/userconfig"
	"github.com/tellmeac/ExtendedSchedule/userconfig/infrastructure/ent"
	uc "github.com/tellmeac/ExtendedSchedule/userconfig/infrastructure/ent/userconfig"
	errs "github.com/tellmeac/ExtendedSchedule/userconfig/middle/errors"
)

// UserConfigDAO provides methods to access UserConfig object.
type UserConfigDAO interface {
	// GetByEmail receives user config by email.
	GetByEmail(ctx context.Context, email string) (*userconfig.UserConfig, error)
	// Onboard will check if config object exists in table, if not insert.
	Onboard(ctx context.Context, userConfig *userconfig.UserConfig) error
	// Update updates user config to desired state.
	Update(ctx context.Context, desired *userconfig.UserConfig) error
}

// NewUserConfigDAO returns a default implementation of UserConfigDAO.
func NewUserConfigDAO(client *ent.Client) UserConfigDAO {
	return &userConfigDAO{client: client}
}

type userConfigDAO struct {
	client *ent.Client
}

func (d userConfigDAO) GetByEmail(ctx context.Context, email string) (*userconfig.UserConfig, error) {
	userConfig, err := d.client.UserConfig.Query().Where(uc.EmailEqualFold(email)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, fmt.Errorf("user config was not found by email (%q): %w", email, errs.ErrNotFound)
	case err != nil:
		return nil, err
	default:
		return toCommonUserConfig(userConfig), nil
	}
}

func (d userConfigDAO) Onboard(ctx context.Context, userConfig *userconfig.UserConfig) error {
	id, err := d.client.UserConfig.Create().
		SetID(userConfig.ID).
		SetEmail(userConfig.Email).
		SetBaseGroup(userConfig.BaseGroup).
		SetExtendedGroupLessons(userConfig.ExtendedGroupLessons).
		SetExcludeRules(userConfig.ExcludeRules).
		OnConflictColumns("id", "email").
		UpdateNewValues().ID(ctx)

	userConfig.ID = id
	return err
}

func (d userConfigDAO) Update(ctx context.Context, desired *userconfig.UserConfig) error {
	updatedCount, err := d.client.UserConfig.Update().
		Where(uc.IDEQ(desired.ID)).
		SetBaseGroup(desired.BaseGroup).
		SetExcludeRules(desired.ExcludeRules).
		SetExtendedGroupLessons(desired.ExtendedGroupLessons).
		Save(ctx)

	if err != nil {
		return err
	}

	if updatedCount < 0 {
		return fmt.Errorf("user config was not found (id = %d): %w", desired.ID, errs.ErrNotFound)
	}
	return nil
}
