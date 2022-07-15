package dao

import (
	"github.com/tellmeac/ExtendedSchedule/userconfig/domain/userconfig"
	"github.com/tellmeac/ExtendedSchedule/userconfig/infrastructure/ent"
)

// toCommonUserConfig converts database object to common user config model.
func toCommonUserConfig(c *ent.UserConfig) *userconfig.UserConfig {
	return &userconfig.UserConfig{
		ID:                   c.ID,
		Email:                c.Email,
		BaseGroup:            c.BaseGroup,
		ExtendedGroupLessons: c.ExtendedGroupLessons,
		ExcludeRules:         c.ExcludeRules,
	}
}
