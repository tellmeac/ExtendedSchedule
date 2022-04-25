package entity

import (
	"github.com/google/uuid"
)

type ExcludedLesson struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	LessonRef  LessonRef
	ByTeacher  *TeacherInfo
	ByPosition int
	ByWeekDays []int
}
