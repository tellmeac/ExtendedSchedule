package entity

import (
	"github.com/google/uuid"
	"time"
)

type ExcludedLesson struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	LessonRef ExcludedRef
}

type ExcludedRef struct {
	LessonID   string
	ByTeacher  *TeacherInfo
	ByWeekDays []time.Weekday
}
