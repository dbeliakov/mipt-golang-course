package students

type DocNumber string

type Student struct {
	FirstName string
	LastName  string
	Grades    []float64
}

func (s *Student) AddGrade(grade float64) {
}

func (s *Student) CalculateAverage() float64 {
	return 0
}

type University struct {
}

func NewUniversity() *University {
	return nil
}

func (u *University) AddStudent(num DocNumber, student Student) {
}

func (u *University) GetStudent(num DocNumber) *Student {
	return nil
}

func (u *University) RemoveStudent(num DocNumber) {
}

func (u *University) CalculateAverage() float64 {
	return 0
}

func (u *University) StudentsSortedByAverageGrade() []Student {
	return nil
}
