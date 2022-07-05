package lesson

// Kind represents lesson type enum.
type Kind string

func (lt Kind) String() string {
	return string(lt)
}

const (
	Lecture      Kind = "LECTURE"
	Practice     Kind = "PRACTICE"
	Seminar      Kind = "SEMINAR"
	Laboratory   Kind = "LABORATORY"
	Exam         Kind = "EXAM"
	DiffCredit   Kind = "DIFFERENTIAL_CREDIT"
	Consultation Kind = "CONSULTATION"
)
