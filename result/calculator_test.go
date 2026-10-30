package result

import (
	"testing"

	"student-result-cli/models"
)

func TestValidateMarks(t *testing.T) {
	tests := []struct {
		name    string
		mark    float64
		wantErr bool
	}{
		{"Valid minimum mark", 0.0, false},
		{"Valid passing mark", 50.0, false},
		{"Valid maximum mark", 100.0, false},
		{"Valid arbitrary mark", 87.5, false},
		{"Invalid negative mark", -1.0, true},
		{"Invalid over maximum mark", 100.5, true},
		{"Invalid extreme mark", 250.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateMarks(tt.mark)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateMarks(%v) error = %v, wantErr %v", tt.mark, err, tt.wantErr)
			}
		})
	}
}

func TestCalculateTotalAndAverage(t *testing.T) {
	subjectMarks := []models.SubjectMark{
		{SubjectName: "Mathematics", Marks: 85.0},
		{SubjectName: "Physics", Marks: 90.0},
		{SubjectName: "Chemistry", Marks: 75.0},
		{SubjectName: "Computer Science", Marks: 95.0},
		{SubjectName: "English", Marks: 80.0},
	}

	total := CalculateTotal(subjectMarks)
	expectedTotal := 425.0
	if total != expectedTotal {
		t.Errorf("CalculateTotal() = %v, expected %v", total, expectedTotal)
	}

	average := CalculateAverage(total, len(subjectMarks))
	expectedAvg := 85.0
	if average != expectedAvg {
		t.Errorf("CalculateAverage() = %v, expected %v", average, expectedAvg)
	}

	// Test zero count edge case
	zeroAvg := CalculateAverage(100.0, 0)
	if zeroAvg != 0.0 {
		t.Errorf("CalculateAverage with 0 count = %v, expected 0.0", zeroAvg)
	}
}

func TestAssignGrade(t *testing.T) {
	tests := []struct {
		name             string
		average          float64
		hasFailedSubject bool
		expectedGrade    string
	}{
		{"Grade A+ for 95", 95.0, false, "A+"},
		{"Grade A+ boundary 90", 90.0, false, "A+"},
		{"Grade A boundary 89.9", 89.9, false, "A"},
		{"Grade A for 80", 80.0, false, "A"},
		{"Grade B for 75", 75.0, false, "B"},
		{"Grade C for 65", 65.0, false, "C"},
		{"Grade D for 55", 55.0, false, "D"},
		{"Grade F for 45", 45.0, false, "F"},
		{"Grade F when a subject is failed despite high average", 92.0, true, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grade := AssignGrade(tt.average, tt.hasFailedSubject)
			if grade != tt.expectedGrade {
				t.Errorf("AssignGrade(%v, %v) = %v, want %v", tt.average, tt.hasFailedSubject, grade, tt.expectedGrade)
			}
		})
	}
}

func TestProcessStudentResultPass(t *testing.T) {
	student := models.Student{
		Name:       "Ezhilarasi K",
		RollNumber: "21CS101",
		SubjectMarks: []models.SubjectMark{
			{SubjectName: "Subject 1", Marks: 92.0},
			{SubjectName: "Subject 2", Marks: 88.0},
			{SubjectName: "Subject 3", Marks: 95.0},
			{SubjectName: "Subject 4", Marks: 91.0},
			{SubjectName: "Subject 5", Marks: 84.0},
		},
	}

	ProcessStudentResult(&student)

	expectedTotal := 450.0
	expectedAverage := 90.0
	expectedGrade := "A+"

	if student.TotalMarks != expectedTotal {
		t.Errorf("TotalMarks = %v, want %v", student.TotalMarks, expectedTotal)
	}
	if student.AverageMarks != expectedAverage {
		t.Errorf("AverageMarks = %v, want %v", student.AverageMarks, expectedAverage)
	}
	if student.Grade != expectedGrade {
		t.Errorf("Grade = %v, want %v", student.Grade, expectedGrade)
	}
	if !student.IsPass {
		t.Errorf("IsPass = %v, want true", student.IsPass)
	}
}

func TestProcessStudentResultFailSubject(t *testing.T) {
	student := models.Student{
		Name:       "Rahul Sharma",
		RollNumber: "21CS102",
		SubjectMarks: []models.SubjectMark{
			{SubjectName: "Subject 1", Marks: 95.0},
			{SubjectName: "Subject 2", Marks: 92.0},
			{SubjectName: "Subject 3", Marks: 45.0}, // Failed subject (<50)
			{SubjectName: "Subject 4", Marks: 88.0},
			{SubjectName: "Subject 5", Marks: 90.0},
		},
	}

	ProcessStudentResult(&student)

	if student.IsPass {
		t.Errorf("Expected IsPass = false because student failed Subject 3, got true")
	}
	if student.Grade != "F" {
		t.Errorf("Expected Grade = 'F' due to failed subject, got %s", student.Grade)
	}
}
