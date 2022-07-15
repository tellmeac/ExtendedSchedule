package userconfig

import (
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/userconfig/domain/values"
)

// NewUserConfig creates clear new UserConfig.
func NewUserConfig(email string) UserConfig {
	return UserConfig{
		ID:    uuid.New(),
		Email: email,
	}
}

// UserConfig represents user configuration.
type UserConfig struct {
	ID                   uuid.UUID              `json:"id"`
	Email                string                 `json:"email"`
	BaseGroup            *values.StudyGroup     `json:"baseGroup"`
	ExcludeRules         []ExcludeRule          `json:"excludeRules"`
	ExtendedGroupLessons []ExtendedGroupLessons `json:"extendedGroupLessons"`
}

// ExtendedGroupLessons represents extended group by some lessons.
type ExtendedGroupLessons struct {
	Group     values.StudyGroup `json:"group"`
	LessonIDs []string          `json:"lessonIds"`
}

// ExcludeRule represents single rule to ignore one or many lessons from schedule repeatedly.
type ExcludeRule struct {
	ID       uuid.UUID           `json:"id"`
	LessonID string              `json:"lessonId"`
	WeekDay  int                 `json:"weekDay"`
	Position int                 `json:"position"`
	Groups   []values.StudyGroup `json:"groups"`
}
