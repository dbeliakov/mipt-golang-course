package courses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessApplications(t *testing.T) {
	testCases := []struct {
		name     string
		students []StudentApplication
		courses  []Course
		expected map[CourseName][]StudentName
	}{
		{
			name: "Single student, single course, within limit",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math": {{FirstName: "John", LastName: "Doe"}},
			},
		},
		{
			name: "Single student, single course, over limit",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 0},
			},
			expected: map[CourseName][]StudentName{},
		},
		{
			name: "Multiple students, single course, within limit",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Math"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
			},
			expected: map[CourseName][]StudentName{
				"Math": {
					{FirstName: "John", LastName: "Doe"},
					{FirstName: "Jane", LastName: "Smith"},
				},
			},
		},
		{
			name: "Multiple students, single course, over limit",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Math"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math": {{FirstName: "John", LastName: "Doe"}},
			},
		},
		{
			name: "Multiple students, multiple courses, within limit",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Physics", "Math"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Multiple students, multiple courses, over limit",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Alice", LastName: "Johnson"}, Avg: 80, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Physics", "Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with same average, different names",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 90, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
			},
			expected: map[CourseName][]StudentName{
				"Math": {
					{FirstName: "John", LastName: "Doe"},
					{FirstName: "Jane", LastName: "Smith"},
				},
			},
		},
		{
			name: "Students with same average and last name, different first names",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
			},
			expected: map[CourseName][]StudentName{
				"Math": {
					{FirstName: "Jane", LastName: "Doe"},
					{FirstName: "John", LastName: "Doe"},
				},
			},
		},
		{
			name: "Students with different averages, same names",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 85, Priorities: []CourseName{"Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
			},
			expected: map[CourseName][]StudentName{
				"Math": {
					{FirstName: "John", LastName: "Doe"},
					{FirstName: "John", LastName: "Doe"},
				},
			},
		},
		{
			name: "Students with no priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
			},
			expected: map[CourseName][]StudentName{},
		},
		{
			name:     "Courses with no students",
			students: []StudentApplication{},
			courses: []Course{
				{Name: "Math", Limit: 2},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{},
		},
		{
			name: "Students with non-existent course priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Chemistry"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Biology"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{},
		},
		{
			name: "Students with multiple priorities, first priority full",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "Alice", LastName: "Johnson"}, Avg: 80, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Math", "Physics"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 2},
			},
			expected: map[CourseName][]StudentName{
				"Math": {{FirstName: "John", LastName: "Doe"}},
				"Physics": {
					{FirstName: "Alice", LastName: "Johnson"},
					{FirstName: "Jane", LastName: "Smith"},
				},
			},
		},
		{
			name: "Students with multiple priorities, second priority full",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Physics", "Math"}},
				{Name: StudentName{FirstName: "Alice", LastName: "Johnson"}, Avg: 80, Priorities: []CourseName{"Physics", "Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}, {FirstName: "Alice", LastName: "Johnson"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with multiple priorities, all priorities full",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Physics", "Math"}},
				{Name: StudentName{FirstName: "Alice", LastName: "Johnson"}, Avg: 80, Priorities: []CourseName{"Math", "Physics"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with same average, different priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 90, Priorities: []CourseName{"Physics", "Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with same average, same priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with different averages, same priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Math", "Physics"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with different averages, different priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 85, Priorities: []CourseName{"Physics", "Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "Jane", LastName: "Smith"}},
			},
		},
		{
			name: "Students with same average, same names, different priorities",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Math", "Physics"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 90, Priorities: []CourseName{"Physics", "Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 1},
				{Name: "Physics", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math":    {{FirstName: "John", LastName: "Doe"}},
				"Physics": {{FirstName: "John", LastName: "Doe"}},
			},
		},
		{
			name: "Big",
			students: []StudentApplication{
				{Name: StudentName{FirstName: "Bob", LastName: "Brown"}, Avg: 88, Priorities: []CourseName{"Biology", "Chemistry", "Physics"}},
				{Name: StudentName{FirstName: "Eve", LastName: "Garcia"}, Avg: 80, Priorities: []CourseName{"Biology", "Physics", "Chemistry"}},
				{Name: StudentName{FirstName: "Hank", LastName: "Jones"}, Avg: 75, Priorities: []CourseName{"Physics", "Math", "Chemistry", "History"}},
				{Name: StudentName{FirstName: "John", LastName: "Doe"}, Avg: 95, Priorities: []CourseName{"Math", "Physics", "Chemistry"}},
				{Name: StudentName{FirstName: "Charlie", LastName: "Davis"}, Avg: 85, Priorities: []CourseName{"Math", "Biology", "Chemistry"}},
				{Name: StudentName{FirstName: "Grace", LastName: "Jones"}, Avg: 75, Priorities: []CourseName{"Math", "Physics", "Biology", "History"}},
				{Name: StudentName{FirstName: "Alice", LastName: "Johnson"}, Avg: 88, Priorities: []CourseName{"Chemistry", "Physics", "Math"}},
				{Name: StudentName{FirstName: "Jane", LastName: "Smith"}, Avg: 95, Priorities: []CourseName{"Physics", "Math", "Biology"}},
				{Name: StudentName{FirstName: "Frank", LastName: "Harris"}, Avg: 78, Priorities: []CourseName{"Chemistry", "Biology", "Physics"}},
				{Name: StudentName{FirstName: "Diana", LastName: "Evans"}, Avg: 83, Priorities: []CourseName{"Physics", "Chemistry", "Math"}},
			},
			courses: []Course{
				{Name: "Math", Limit: 2},
				{Name: "Physics", Limit: 2},
				{Name: "Chemistry", Limit: 2},
				{Name: "Biology", Limit: 2},
				{Name: "History", Limit: 1},
			},
			expected: map[CourseName][]StudentName{
				"Math": {
					{FirstName: "Charlie", LastName: "Davis"},
					{FirstName: "John", LastName: "Doe"},
				},
				"Physics": {
					{FirstName: "Diana", LastName: "Evans"},
					{FirstName: "Jane", LastName: "Smith"},
				},
				"Chemistry": {
					{FirstName: "Frank", LastName: "Harris"},
					{FirstName: "Alice", LastName: "Johnson"},
				},
				"Biology": {
					{FirstName: "Bob", LastName: "Brown"},
					{FirstName: "Eve", LastName: "Garcia"},
				},
				"History": {
					{FirstName: "Grace", LastName: "Jones"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ProcessApplications(tc.students, tc.courses)
			assert.Equal(t, tc.expected, result)
		})
	}
}
