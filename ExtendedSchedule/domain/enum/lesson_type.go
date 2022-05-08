package enum

type LessonType string

const (
	Lecture    LessonType = "LECTURE"
	Practice   LessonType = "PRACTICE"
	Seminar    LessonType = "SEMINAR"
	Laboratory LessonType = "LABORATORY"
	Exam       LessonType = "EXAM"
)
