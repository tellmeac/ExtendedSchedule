package aggregate

import (
	"tellmeac/extended-schedule/domain/entity"

	"github.com/google/uuid"
)

func NewUserConfig(email string) UserConfig {
	return UserConfig{
		ID:    uuid.New(),
		Email: email,
	}
}

// UserConfig представляет объект конфигурации пользователя.
type UserConfig struct {
	ID                   uuid.UUID               `json:"id"`
	Email                string                  `json:"email"`
	BaseGroup            *entity.GroupInfo       `json:"baseGroup"`
	ExcludedLessons      []entity.ExcludedLesson `json:"excludedLessons"`
	ExtendedGroupLessons []ExtendedGroupLessons  `json:"extendedGroupLessons"`
}

// ExtendedGroupLessons представляет группу с дополнительными предметами.
type ExtendedGroupLessons struct {
	Group     entity.GroupInfo `json:"group"`
	LessonIDs []string         `json:"lessonIds"`
}
