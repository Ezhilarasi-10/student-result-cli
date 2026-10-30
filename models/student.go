package models

// Constants representing system thresholds and constraints.
const (
	// MinValidMark represents the lowest possible score for any subject.
	MinValidMark float64 = 0.0

	// MaxValidMark represents the highest possible score for any subject.
	MaxValidMark float64 = 100.0

	// PassingMarkThreshold represents the minimum score required in each subject to pass.
	PassingMarkThreshold float64 = 50.0

	// DefaultSubjectCount specifies the standard number of subjects evaluated.
	DefaultSubjectCount int = 5
)

// SubjectMark holds the name of a subject and the marks obtained by a student.
type SubjectMark struct {
	SubjectName string  `json:"subject_name"`
	Marks       float64 `json:"marks"`
}

// Student holds personal information and academic evaluation results for a student.
type Student struct {
	Name         string        `json:"name"`
	RollNumber   string        `json:"roll_number"`
	SubjectMarks []SubjectMark `json:"subject_marks"`
	TotalMarks   float64       `json:"total_marks"`
	AverageMarks float64       `json:"average_marks"`
	Grade        string        `json:"grade"`
	IsPass       bool          `json:"is_pass"`
}
