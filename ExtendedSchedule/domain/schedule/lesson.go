package schedule

import "tellmeac/extended-schedule/domain/values"

// Lesson represents lesson.
type Lesson struct {
	ID         string             `json:"id"`
	Title      string             `json:"title"`
	LessonType LessonType         `json:"lessonType"`
	Teacher    values.TeacherInfo `json:"teacher"`
	Groups     []values.GroupInfo `json:"groups"`
}

// LessonInSchedule represents lesson in the context of one day of the schedule.
type LessonInSchedule struct {
	ID         string              `json:"id"`
	Title      string              `json:"title"`
	Position   int                 `json:"position"`
	LessonType LessonType          `json:"lessonType"`
	Teacher    values.TeacherInfo  `json:"teacher"`
	Audience   values.AudienceInfo `json:"audience"`
	Groups     []values.GroupInfo  `json:"groups"`
}

// IsDuplicate returns true if lesson is duplicated in context.
func (lesson LessonInSchedule) IsDuplicate(other *LessonInSchedule) bool {
	return lesson.ID == other.ID &&
		lesson.Position == other.Position &&
		lesson.Teacher == other.Teacher &&
		lesson.LessonType == other.LessonType
}
