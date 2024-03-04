package weekday

import "testing"

func TestNextDay(t *testing.T) {
	testCases := []struct {
		name     string
		current  Weekday
		expected Weekday
	}{
		{"From Monday to Tuesday", Monday, Tuesday},
		{"From Sunday to Monday", Sunday, Monday},
		{"From Friday to Saturday", Friday, Saturday},
		{"From Thursday to Friday", Thursday, Friday},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := NextDay(testCase.current)
			if actual != testCase.expected {
				t.Errorf("Expected %d, got %d", testCase.expected, actual)
			}
		})
	}
}

func TestNextDayCycle(t *testing.T) {
	days := []Weekday{Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday}
	for i, day := range days {
		next := NextDay(day)
		expected := days[(i+1)%len(days)]
		if next != expected {
			t.Errorf("After %d, expected %d, got %d", day, expected, next)
		}
	}
}
