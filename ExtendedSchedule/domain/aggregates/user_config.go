package aggregates

import (
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/entity"
)

type UserConfig struct {
	UserID          uuid.UUID
	JoinedGroups    []entity.GroupInfo
	ExcludedLessons []entity.ExcludedLesson
}
