package enum

type LessonType string

const (
	Lecture      LessonType = "LECTURE"
	Practice     LessonType = "PRACTICE"
	Seminar      LessonType = "SEMINAR"
	Laboratory   LessonType = "LABORATORY"
	Exam         LessonType = "EXAM"
	DiffCredit   LessonType = "DIFFERENTIAL_CREDIT"
	Consultation LessonType = "CONSULTATION"
)
