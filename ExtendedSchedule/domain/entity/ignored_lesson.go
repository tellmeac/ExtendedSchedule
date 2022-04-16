package entity

import "github.com/google/uuid"

type ExcludedLesson struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	LessonRef uuid.UUID
}
