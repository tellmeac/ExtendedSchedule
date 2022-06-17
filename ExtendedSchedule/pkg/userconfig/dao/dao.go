package dao

import (
	"context"
	"fmt"
	commonmodels "tellmeac/extended-schedule/common/models"
	errs "tellmeac/extended-schedule/lib/errors"
	"tellmeac/extended-schedule/pkg/ent"
	"tellmeac/extended-schedule/pkg/ent/userconfig"
)

// DAO provides methods to access UserConfig object.
type DAO interface {
	// GetByEmail receives user config by email.
	GetByEmail(ctx context.Context, email string) (*commonmodels.UserConfig, error)
	// Onboard will check if config object exists in table, if not insert.
	Onboard(ctx context.Context, userConfig *commonmodels.UserConfig) error
	// Update updates user config to desired state.
	Update(ctx context.Context, desired *commonmodels.UserConfig) error
}

// New returns a default implementation of DAO.
func New(client *ent.Client) DAO {
	return &dao{client: client}
}

type dao struct {
	client *ent.Client
}

func (d dao) GetByEmail(ctx context.Context, email string) (*commonmodels.UserConfig, error) {
	userConfig, err := d.client.UserConfig.Query().Where(userconfig.EmailEqualFold(email)).Only(ctx)
	switch {
	case ent.IsNotFound(err):
		return nil, fmt.Errorf("user config was not found by email (%q): %w", email, errs.ErrNotFound)
	case err != nil:
		return nil, err
	default:
		return toCommonUserConfig(userConfig), nil
	}
}

func (d dao) Onboard(ctx context.Context, userConfig *commonmodels.UserConfig) error {
	id, err := d.client.UserConfig.Create().
		SetID(userConfig.ID).
		SetEmail(userConfig.Email).
		SetBaseGroup(userConfig.BaseGroup).
		SetExtendedGroupLessons(userConfig.ExtendedGroupLessons).
		SetExcludeRules(userConfig.ExcludedLessons).
		OnConflictColumns("id", "email").
		UpdateNewValues().ID(ctx)

	userConfig.ID = id
	return err
}

func (d dao) Update(ctx context.Context, desired *commonmodels.UserConfig) error {
	updatedCount, err := d.client.UserConfig.Update().
		Where(userconfig.IDEQ(desired.ID)).
		SetBaseGroup(desired.BaseGroup).
		SetExcludeRules(desired.ExcludedLessons).
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
