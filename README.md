# Student Result Processing CLI Application

> **College Go Programming Lab - Assignment I**  
> **Course / Subject:** Go (Golang) Programming  
> **Author:** College Student Project  
> **Version:** 1.0.0

---

## Table of Contents
1. [Introduction](#introduction)
2. [Objectives](#objectives)
3. [Technologies Used](#technologies-used)
4. [Features](#features)
5. [Folder Structure](#folder-structure)
6. [Grading System](#grading-system)
7. [Installation & Prerequisites](#installation--prerequisites)
8. [How to Run the Application](#how-to-run-the-application)
9. [Running Unit Tests](#running-unit-tests)
10. [Sample Terminal Session (Input & Output)](#sample-terminal-session-input--output)
11. [Git Version Control & GitHub Upload Instructions](#git-version-control--github-upload-instructions)

---

## Introduction
The **Student Result Processing CLI** is a modular, command-line tool built using **Go (Golang)**. It allows educators and students to enter student details, record marks across 5 subjects, validate scores, calculate total marks and percentage average, evaluate pass/fail criteria, and generate a formatted academic report card right in the terminal.

This project was built following industry-standard Go best practices, utilizing Go modules, custom packages (`models` and `result`), unit testing, and Git version control with feature branches.

---

## Objectives
- Demonstrate foundational Go concepts:
  - Variables, typed constants, and standard data types (`string`, `float64`, `int`, `bool`, slices, and structs).
  - Formatted I/O using the standard `fmt` and `bufio` packages.
  - Control flow using loops (`for`) and conditional branching (`if-else`, `switch`).
- Implement robust input validation (ensuring marks fall between `0` and `100`).
- Process results using modular business logic:
  - Calculate total marks and arithmetic average.
  - Assign grades from `A+` to `F`.
  - Enforce subject-level pass criteria (if any subject mark is below 50, the student is marked as **FAIL**).
- Organize code into clean, reusable packages.
- Implement unit tests for key calculation routines.
- Maintain version control using Git with feature branching.

---

## Technologies Used
- **Language:** Go (Golang) 1.22+
- **Standard Libraries:**
  - `fmt` (Formatted terminal I/O)
  - `bufio` & `os` (Buffered terminal input)
  - `strconv` & `strings` (Data parsing and string manipulation)
  - `errors` (Custom error management)
  - `testing` (Go standard testing framework)
- **Version Control:** Git & GitHub

---

## Features
- **Interactive Terminal Interface:** Clean prompts guiding the user to enter student name, roll number, and 5 subject marks.
- **Strict Input Validation:** Prevents negative numbers, scores above 100, and non-numeric inputs by re-prompting with clear error messages.
- **Fail-Safe Academic Rules:** Any subject score below 50 automatically triggers a `FAIL` status and an `F` grade, regardless of how high the average score might be.
- **Formatted Report Card:** Prints a clean ASCII performance table summarizing subject marks, total score, percentage, grade, and final result.
- **Looping Workflow:** Prompt to process another student or safely exit.

---

## Folder Structure
```
student-result-cli/
├── go.mod                     # Go module definition
├── main.go                    # CLI entrypoint and user interaction loop
├── models/
│   └── student.go             # Data models (Student, SubjectMark) and constants
├── result/
│   ├── calculator.go          # Core calculation, grading, validation, and display logic
│   └── calculator_test.go     # Unit tests covering all calculation edge cases
└── README.md                  # Comprehensive documentation
```

---

## Grading System

| Marks Range (Average) | Grade | Status | Notes |
|:----------------------|:-----:|:------:|:-----------------------------------------------|
| 90 – 100              | **A+** | PASS   | Requires passing mark (>= 50) in all 5 subjects |
| 80 – 89.99            | **A**  | PASS   | Requires passing mark (>= 50) in all 5 subjects |
| 70 – 79.99            | **B**  | PASS   | Requires passing mark (>= 50) in all 5 subjects |
| 60 – 69.99            | **C**  | PASS   | Requires passing mark (>= 50) in all 5 subjects |
| 50 – 59.99            | **D**  | PASS   | Requires passing mark (>= 50) in all 5 subjects |
| Below 50              | **F**  | FAIL   | Failed overall average |
| *Any subject < 50*    | **F**  | FAIL   | Mandatory fail rule applied |

---

## Installation & Prerequisites

1. Ensure Go 1.22 or newer is installed:
   ```bash
   go version
   ```
2. Navigate to the `student-result-cli` directory:
   ```bash
   cd student-result-cli
   ```
3. Verify that dependencies and packages are organized:
   ```bash
   go mod tidy
   ```

---

## How to Run the Application

### Option 1: Run directly with `go run`
```bash
go run main.go
```

### Option 2: Build an executable binary
```bash
# Build the executable
go build -o student-result-cli.exe .

# Execute on Windows:
.\student-result-cli.exe

# Execute on Linux / macOS:
./student-result-cli
```

---

## Running Unit Tests
Execute the unit tests to verify calculation logic and grade assignments:
```bash
go test -v ./...
```
Expected output:
```
=== RUN   TestValidateMarks
--- PASS: TestValidateMarks (0.00s)
=== RUN   TestCalculateTotalAndAverage
--- PASS: TestCalculateTotalAndAverage (0.00s)
=== RUN   TestAssignGrade
--- PASS: TestAssignGrade (0.00s)
=== RUN   TestProcessStudentResultPass
--- PASS: TestProcessStudentResultPass (0.00s)
=== RUN   TestProcessStudentResultFailSubject
--- PASS: TestProcessStudentResultFailSubject (0.00s)
PASS
ok      student-result-cli/result       0.045s
```

---

## Sample Terminal Session (Input & Output)

```text
============================================================
         STUDENT RESULT PROCESSING SYSTEM (CLI)
============================================================
Enter Student Name: Ezhilarasi K
Enter Roll Number: 21CS101

Enter marks for 5 subjects (Valid range: 0 - 100):
  Enter marks for Mathematics (0-100): 92
  Enter marks for Data Structures (0-100): 88
  Enter marks for Operating Systems (0-100): 95
  Enter marks for Database Management (0-100): 91
  Enter marks for Computer Networks (0-100): 84

============================================================
             STUDENT ACADEMIC PERFORMANCE REPORT            
============================================================
Student Name : Ezhilarasi K                  
Roll Number  : 21CS101                       
------------------------------------------------------------
S.No  | Subject Name                 | Marks      | Status
------------------------------------------------------------
1     | Mathematics                  |      92.00 | PASS  
2     | Data Structures              |      88.00 | PASS  
3     | Operating Systems            |      95.00 | PASS  
4     | Database Management          |      91.00 | PASS  
5     | Computer Networks            |      84.00 | PASS  
------------------------------------------------------------
Total Marks  :     450.00 / 500.00
Average      :      90.00%
Grade        : A+        
Final Result : PASSED
============================================================

Do you want to process another student result? (y/n): n

Thank you for using the Student Result Processing System. Goodbye!
```

---

## Git Version Control & GitHub Upload Instructions

This repository is structured to be independently pushed to GitHub:

### 1. Initialize Git and Stage Files
```bash
git init
git add .
git commit -m "feat: initial commit for student result processing CLI"
```

### 2. Feature Branching Workflow
```bash
# Create and switch to a feature branch
git checkout -b feature/grading-system

# Make improvements or edits, then commit
git commit -am "feat: implement grade calculation and validation logic"

# Switch back to main and merge
git checkout main
git merge feature/grading-system
```

### 3. Push to GitHub
1. Create a new repository on [GitHub](https://github.com/new) named `student-result-cli`.
2. Connect your local repository and push:
```bash
git remote add origin https://github.com/<your-username>/student-result-cli.git
git branch -M main
git push -u origin main
```
