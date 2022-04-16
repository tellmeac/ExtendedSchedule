package entity

import "github.com/google/uuid"

type IgnoredLesson struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	LessonRef uuid.UUID
}
