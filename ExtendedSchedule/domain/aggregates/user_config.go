package aggregates

import (
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/entity"
)

type UserConfig struct {
	UserID          uuid.UUID               `json:"userID"`
	JoinedGroups    []entity.GroupInfo      `json:"joinedGroups"`
	ExcludedLessons []entity.ExcludedLesson `json:"excludedLessons"`
}
