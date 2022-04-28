package aggregates

import (
	"github.com/google/uuid"
	"tellmeac/extended-schedule/domain/entity"
)

type UserConfig struct {
	UserID          uuid.UUID               `json:"UserID"`
	JoinedGroups    []entity.GroupInfo      `json:"JoinedGroups"`
	ExcludedLessons []entity.ExcludedLesson `json:"ExcludedLessons"`
}
