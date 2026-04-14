package courses

type StudentName struct {
	FirstName string
	LastName  string
}

func (s *StudentName) String() string {
	return s.LastName + " " + s.FirstName
}

type CourseName string

type StudentApplication struct {
	Name       StudentName
	Avg        uint32
	Priorities []CourseName
}

type Course struct {
	Name  CourseName
	Limit uint32
}

func ProcessApplications(students []StudentApplication, courses []Course) (map[CourseName][]StudentName, []StudentApplication) {
	panic("implement me")
}
