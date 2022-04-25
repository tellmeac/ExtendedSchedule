package entity

const (
	EmptyLesson = "EMPTY"
)

type Lesson struct {
	ID         string
	Title      string
	LessonType string
	Teacher    TeacherInfo
	Audience   AudienceInfo
	Groups     []GroupInfo
}

type LessonRef struct {
	LessonID string
}

type AudienceInfo struct {
	ID   string
	Name string
}

type GroupInfo struct {
	ID   string
	Name string
}

type TeacherInfo struct {
	ID   string
	Name string
}
