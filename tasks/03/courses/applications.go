package courses

type StudentName struct {
	FirstName string
	LastName  string
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

func ProcessApplications(students []StudentApplication, courses []Course) map[CourseName][]StudentName {
	panic("implement me")
}
