package entity

import (
	"github.com/google/uuid"
)

type ExcludedLesson struct {
	ID       uuid.UUID
	LessonID string
	Teacher  *TeacherInfo
	Position int
	WeekDay  int
}
