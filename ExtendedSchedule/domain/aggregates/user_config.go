package aggregates

import (
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/entity"
)

type UserConfig struct {
	UserID            uuid.UUID
	JoinedGroups      []entity.GroupInfo
	IgnoredLessons    []entity.ExcludedLesson
	AdditionalLessons []entity.ExtendedLesson
}
