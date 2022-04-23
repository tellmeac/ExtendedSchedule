package entity

import (
	"github.com/google/uuid"
)

type ExtendedLesson struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Description string
	Ref         LessonRef
}

type LessonRef struct {
	LessonID string
	GroupID  string
}
