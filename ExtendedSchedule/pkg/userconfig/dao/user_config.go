package dao

import (
	commonmodels "tellmeac/extended-schedule/common/models"
	"tellmeac/extended-schedule/pkg/ent"
)

// toCommonUserConfig converts database object to common user config model.
func toCommonUserConfig(c *ent.UserConfig) *commonmodels.UserConfig {
	return &commonmodels.UserConfig{
		ID:                   c.ID,
		Email:                c.Email,
		BaseGroup:            c.BaseGroup,
		ExtendedGroupLessons: c.ExtendedGroupLessons,
		ExcludedLessons:      c.ExcludeRules,
	}
}
