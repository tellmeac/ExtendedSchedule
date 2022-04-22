package entity

import (
	"tellmeac/extended-schedule/domain/values"

	"github.com/google/uuid"
)

type ExtendedLesson struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Context   values.LessonContext
	Intervals []values.LessonInterval
}
