package models

// Lesson represents lesson.
type Lesson struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	LessonType LessonType  `json:"lessonType"`
	Teacher    TeacherInfo `json:"teacher"`
	Groups     []GroupInfo `json:"groups"`
}

// LessonWithContext represents lesson in the context of one day of the schedule.
type LessonWithContext struct {
	ID         string       `json:"id"`
	Title      string       `json:"title"`
	Position   int          `json:"position"`
	LessonType LessonType   `json:"lessonType"`
	Teacher    TeacherInfo  `json:"teacher"`
	Audience   AudienceInfo `json:"audience"`
	Groups     []GroupInfo  `json:"groups"`
}

// IsDuplicate returns true if lesson is duplicated in context.
func (lesson LessonWithContext) IsDuplicate(other *LessonWithContext) bool {
	return lesson.ID == other.ID &&
		lesson.Position == other.Position &&
		lesson.Teacher == other.Teacher &&
		lesson.LessonType == other.LessonType
}

// LessonType represents lesson type enum.
type LessonType string

func (lt LessonType) String() string {
	return string(lt)
}

const (
	Lecture      LessonType = "LECTURE"
	Practice     LessonType = "PRACTICE"
	Seminar      LessonType = "SEMINAR"
	Laboratory   LessonType = "LABORATORY"
	Exam         LessonType = "EXAM"
	DiffCredit   LessonType = "DIFFERENTIAL_CREDIT"
	Consultation LessonType = "CONSULTATION"
)

type AudienceInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeacherInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FacultyInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
