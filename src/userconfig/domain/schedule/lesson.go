package schedule

import "github.com/tellmeac/ExtendedSchedule/userconfig/domain/lesson"

type Lesson struct {
	lesson.Lesson
	Position int
}

// IsSame returns true if lesson is duplicated in context.
func (lesson Lesson) IsSame(other *Lesson) bool {
	return lesson.ID == other.ID &&
		lesson.Position == other.Position &&
		lesson.Teacher == other.Teacher &&
		lesson.Kind == other.Kind
}
