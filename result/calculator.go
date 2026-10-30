package result

import (
	"errors"
	"fmt"
	"strings"

	"student-result-cli/models"
)

// ErrInvalidMark is returned when a subject mark falls outside the allowable 0-100 range.
var ErrInvalidMark = errors.New("marks must be between 0 and 100")

// ValidateMarks checks whether a given mark is within the valid academic score boundary [0, 100].
func ValidateMarks(mark float64) error {
	if mark < models.MinValidMark || mark > models.MaxValidMark {
		return fmt.Errorf("%w: received %.2f", ErrInvalidMark, mark)
	}
	return nil
}

// CalculateTotal calculates the sum of marks obtained across all subjects.
func CalculateTotal(subjectMarks []models.SubjectMark) float64 {
	var total float64 = 0.0
	for _, sm := range subjectMarks {
		total += sm.Marks
	}
	return total
}

// CalculateAverage calculates the arithmetic mean of the total marks given the subject count.
func CalculateAverage(total float64, count int) float64 {
	if count <= 0 {
		return 0.0
	}
	return total / float64(count)
}

// HasFailedAnySubject checks if the student scored below the passing threshold in any subject.
func HasFailedAnySubject(subjectMarks []models.SubjectMark) bool {
	for _, sm := range subjectMarks {
		if sm.Marks < models.PassingMarkThreshold {
			return true
		}
	}
	return false
}

// AssignGrade assigns a letter grade based on the student's average and subject pass status.
// Grading criteria:
//
//	A+ : 90 - 100
//	A  : 80 - 89.99
//	B  : 70 - 79.99
//	C  : 60 - 69.99
//	D  : 50 - 59.99
//	F  : Below 50 (or if any individual subject mark is below 50)
func AssignGrade(average float64, hasFailedSubject bool) string {
	if hasFailedSubject {
		return "F"
	}

	switch {
	case average >= 90.0 && average <= 100.0:
		return "A+"
	case average >= 80.0 && average < 90.0:
		return "A"
	case average >= 70.0 && average < 80.0:
		return "B"
	case average >= 60.0 && average < 70.0:
		return "C"
	case average >= 50.0 && average < 60.0:
		return "D"
	default:
		return "F"
	}
}

// ProcessStudentResult performs complete result processing for a student:
// calculates total marks, average, verifies subject passing criteria, assigns grade, and updates pass/fail status.
func ProcessStudentResult(student *models.Student) {
	student.TotalMarks = CalculateTotal(student.SubjectMarks)
	student.AverageMarks = CalculateAverage(student.TotalMarks, len(student.SubjectMarks))

	failedSubject := HasFailedAnySubject(student.SubjectMarks)
	student.Grade = AssignGrade(student.AverageMarks, failedSubject)

	// A student passes only if they passed every subject AND average is at least 50.
	student.IsPass = !failedSubject && student.AverageMarks >= models.PassingMarkThreshold
}

// DisplayResult prints a formatted, professional terminal report card for the student.
func DisplayResult(student models.Student) {
	divider := strings.Repeat("=", 60)
	subDivider := strings.Repeat("-", 60)

	fmt.Println()
	fmt.Println(divider)
	fmt.Println("             STUDENT ACADEMIC PERFORMANCE REPORT            ")
	fmt.Println(divider)
	fmt.Printf("Student Name : %-30s\n", student.Name)
	fmt.Printf("Roll Number  : %-30s\n", student.RollNumber)
	fmt.Println(subDivider)
	fmt.Printf("%-5s | %-28s | %-10s | %-6s\n", "S.No", "Subject Name", "Marks", "Status")
	fmt.Println(subDivider)

	for i, sm := range student.SubjectMarks {
		status := "PASS"
		if sm.Marks < models.PassingMarkThreshold {
			status = "FAIL"
		}
		fmt.Printf("%-5d | %-28s | %10.2f | %-6s\n", i+1, sm.SubjectName, sm.Marks, status)
	}

	fmt.Println(subDivider)
	fmt.Printf("Total Marks  : %10.2f / %.2f\n", student.TotalMarks, float64(len(student.SubjectMarks))*models.MaxValidMark)
	fmt.Printf("Average      : %10.2f%%\n", student.AverageMarks)
	fmt.Printf("Grade        : %-10s\n", student.Grade)

	resultText := "PASSED"
	if !student.IsPass {
		resultText = "FAILED (Subject mark below 50 or average < 50)"
	}
	fmt.Printf("Final Result : %s\n", resultText)
	fmt.Println(divider)
	fmt.Println()
}
