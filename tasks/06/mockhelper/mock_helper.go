package mockhelper

type TestingT interface {
	Errorf(format string, args ...any)
	FailNow()
	Cleanup(cb func())
}

type Matcher interface {
	Match(any) bool
	String() string
}

type Call interface {
	Return(returns ...any)
}

type MockHelper interface {
	ExpectCall(method string, args ...any) Call
	Call(method string, args ...any) []any
	Finish()
}

func Any() Matcher {
	panic("implement me")
}

func NewMockHelper(t TestingT) MockHelper {
	panic("implement me")
}
