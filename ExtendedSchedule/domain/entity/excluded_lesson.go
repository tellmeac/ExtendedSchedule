package entity

import (
	"github.com/google/uuid"
)

type ExcludedLesson struct {
	ID       uuid.UUID   `json:"id"`
	LessonID string      `json:"lessonId"`
	Position int         `json:"position"`
	WeekDay  int         `json:"weekDay"`
	Groups   []GroupInfo `json:"groups"`
}
