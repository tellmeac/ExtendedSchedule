package aggregates

import (
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

type UserConfig struct {
	UserID            uuid.UUID
	JoinedGroups      []entity.GroupInfo
	IgnoredLessons    []entity.ExcludedLesson
	AdditionalLessons []entity.ExtendedLesson
}
