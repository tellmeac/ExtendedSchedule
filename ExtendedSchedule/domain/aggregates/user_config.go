package aggregates

import (
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/entity"
	"github.com/tellmeac/extended-schedule/domain/values"
)

type UserConfig struct {
	UserID            uuid.UUID
	JoinedGroups      []values.GroupInfo
	IgnoredLessons    []entity.IgnoredLesson
	AdditionalLessons []entity.ExtendedLesson
}
