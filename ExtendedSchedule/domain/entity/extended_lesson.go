package entity

import (
	"github.com/google/uuid"
	"github.com/tellmeac/extended-schedule/domain/values"
)

type ExtendedLesson struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Context   values.LessonContext
	Intervals []values.LessonInterval
}
