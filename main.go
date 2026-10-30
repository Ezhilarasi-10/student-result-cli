package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"student-result-cli/models"
	"student-result-cli/result"
)

// AppTitle defines the application banner title.
const AppTitle = "STUDENT RESULT PROCESSING SYSTEM (CLI)"

// ReadStudentDetails prompts the user via the terminal and gathers student details and marks for 5 subjects.
func ReadStudentDetails(reader *bufio.Reader) (models.Student, error) {
	var student models.Student

	// 1. Read Student Name
	for {
		fmt.Print("Enter Student Name: ")
		nameInput, err := reader.ReadString('\n')
		if err != nil {
			return student, fmt.Errorf("error reading student name: %w", err)
		}
		nameInput = strings.TrimSpace(nameInput)
		if nameInput == "" {
			fmt.Println("Error: Student name cannot be empty. Please enter a valid name.")
			continue
		}
		student.Name = nameInput
		break
	}

	// 2. Read Roll Number
	for {
		fmt.Print("Enter Roll Number: ")
		rollInput, err := reader.ReadString('\n')
		if err != nil {
			return student, fmt.Errorf("error reading roll number: %w", err)
		}
		rollInput = strings.TrimSpace(rollInput)
		if rollInput == "" {
			fmt.Println("Error: Roll number cannot be empty. Please enter a valid roll number.")
			continue
		}
		student.RollNumber = rollInput
		break
	}

	// Default subjects list
	subjectNames := []string{
		"Mathematics",
		"Data Structures",
		"Operating Systems",
		"Database Management",
		"Computer Networks",
	}

	fmt.Printf("\nEnter marks for %d subjects (Valid range: %.0f - %.0f):\n",
		models.DefaultSubjectCount, models.MinValidMark, models.MaxValidMark)

	student.SubjectMarks = make([]models.SubjectMark, 0, models.DefaultSubjectCount)

	// 3. Read Marks for each subject with validation loop
	for i := 0; i < models.DefaultSubjectCount; i++ {
		subjName := subjectNames[i]

		for {
			fmt.Printf("  Enter marks for %s (0-100): ", subjName)
			markInput, err := reader.ReadString('\n')
			if err != nil {
				return student, fmt.Errorf("error reading mark for %s: %w", subjName, err)
			}
			markInput = strings.TrimSpace(markInput)

			markValue, parseErr := strconv.ParseFloat(markInput, 64)
			if parseErr != nil {
				fmt.Println("    [!] Invalid input: Please enter a valid numeric value.")
				continue
			}

			// Validate mark range using result.ValidateMarks
			if valErr := result.ValidateMarks(markValue); valErr != nil {
				fmt.Printf("    [!] Validation Error: %s\n", valErr.Error())
				continue
			}

			student.SubjectMarks = append(student.SubjectMarks, models.SubjectMark{
				SubjectName: subjName,
				Marks:       markValue,
			})
			break
		}
	}

	return student, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("============================================================")
	fmt.Printf("         %s\n", AppTitle)
	fmt.Println("============================================================")

	for {
		student, err := ReadStudentDetails(reader)
		if err != nil {
			// Handle EOF gracefully (useful during automated tests or piping)
			if err.Error() == "EOF" || strings.Contains(err.Error(), "EOF") {
				fmt.Println("\nInput stream closed. Exiting.")
				break
			}
			fmt.Printf("\nAn error occurred while reading details: %v\n", err)
			break
		}

		// Process result: calculate total, average, grade, pass/fail status
		result.ProcessStudentResult(&student)

		// Display formatted result
		result.DisplayResult(student)

		// Ask user if they wish to process another student
		fmt.Print("Do you want to process another student result? (y/n): ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		choice = strings.ToLower(strings.TrimSpace(choice))
		if choice != "y" && choice != "yes" {
			fmt.Println("\nThank you for using the Student Result Processing System. Goodbye!")
			break
		}
		fmt.Println()
	}
}
