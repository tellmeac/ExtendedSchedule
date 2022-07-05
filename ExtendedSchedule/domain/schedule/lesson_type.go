package schedule

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
