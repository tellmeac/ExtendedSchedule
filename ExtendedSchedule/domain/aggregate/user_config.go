package aggregate

import (
	"tellmeac/extended-schedule/domain/entity"
)

type UserConfig struct {
	UserIdentifier  string                  `json:"userIdentifier"`
	JoinedGroups    []entity.GroupInfo      `json:"joinedGroups"`
	ExcludedLessons []entity.ExcludedLesson `json:"excludedLessons"`
}
