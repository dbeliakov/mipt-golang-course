package students

import (
	"math"
	"testing"
)

func TestAddGrade(t *testing.T) {
	s := Student{
		FirstName: "John",
		LastName:  "Doe",
		Grades:    []float64{5.0, 4.0},
	}

	s.AddGrade(3.0)

	if len(s.Grades) != 3 {
		t.Fatalf("Expected 3 grades, got %d", len(s.Grades))
	}

	if s.Grades[2] != 3.0 {
		t.Errorf("Expected last grade to be 3.0, got %f", s.Grades[2])
	}
}

func TestCalculateAverage(t *testing.T) {
	s := Student{
		FirstName: "Jane",
		LastName:  "Doe",
		Grades:    []float64{5.0, 4.0, 3.0},
	}

	avg := s.CalculateAverage()

	if avg != 4.0 {
		t.Errorf("Expected average of 4.0, got %f", avg)
	}
}

func TestUniversity_AddStudentAndGetStudent(t *testing.T) {
	u := NewUniversity()
	student := Student{
		FirstName: "Test",
		LastName:  "Student",
		Grades:    []float64{5.0, 4.0},
	}
	docNumber := DocNumber("123456")

	u.AddStudent(docNumber, student)

	retrievedStudent := u.GetStudent(docNumber)

	if retrievedStudent == nil {
		t.Fatalf("No students found, got nil")
	}

	if retrievedStudent.FirstName != "Test" {
		t.Errorf("Expected first name Test, got %s", retrievedStudent.FirstName)
	}

	if retrievedStudent.LastName != "Student" {
		t.Errorf("Expected last name Student, got %s", retrievedStudent.LastName)
	}
}

func TestUniversity_RemoveStudent(t *testing.T) {
	u := NewUniversity()
	docNumber := DocNumber("123456")
	u.AddStudent(docNumber, Student{})

	u.RemoveStudent(docNumber)

	if u.GetStudent(docNumber) != nil {
		t.Error("Student not removed")
	}
}

func TestUniversity_StudentsSortedByAverageGrade(t *testing.T) {
	u := NewUniversity()
	u.AddStudent("1", Student{
		FirstName: "Alice",
		LastName:  "Zephyr",
		Grades:    []float64{3.0, 3.0},
	})
	u.AddStudent("2", Student{
		FirstName: "Bob",
		LastName:  "Yellow",
		Grades:    []float64{4.0, 4.0},
	})
	u.AddStudent("3", Student{
		FirstName: "Bob",
		LastName:  "Yellow",
		Grades:    []float64{4.0, 3.0},
	})
	sortedStudents := u.StudentsSortedByAverageGrade()

	if len(sortedStudents) != 3 {
		t.Fatalf("Incorrect count of students, got %d", len(sortedStudents))
	}

	if sortedStudents[0].FirstName != "Alice" {
		t.Errorf("Expected Alice to be first, got %s", sortedStudents[0].FirstName)
	}

	if sortedStudents[2].FirstName != "Bob" {
		t.Errorf("Expected Bob to be second, got %s", sortedStudents[1].FirstName)
	}
}

func TestUniversity_CalculateAverage(t *testing.T) {
	u := NewUniversity()
	u.AddStudent("1", Student{
		FirstName: "Alice",
		LastName:  "Zephyr",
		Grades:    []float64{3.0, 4.0},
	})
	u.AddStudent("2", Student{
		FirstName: "Bob",
		LastName:  "Yellow",
		Grades:    []float64{5.0, 4.0},
	})
	u.AddStudent("3", Student{
		FirstName: "Charlie",
		LastName:  "Xenon",
		Grades:    []float64{4.0, 5.0},
	})

	avg := u.CalculateAverage()
	expectedAvg := (3.0 + 4.0 + 5.0 + 4.0 + 4.0 + 5.0) / 6.0

	if math.Abs(avg-expectedAvg) > 0.0001 {
		t.Errorf("Expected university average to be %f, got %f", expectedAvg, avg)
	}
}
