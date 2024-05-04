package httperror

type httpError struct{}

func (e httpError) Error() string {
	return "IMPL ME"
}

var _ error = httpError{}

func NewHTTPError(statusCode int) error {
	return nil
}

func WrapWithHTTPError(err error, statusCode int) error {
	return nil
}

func GetStatusCode(err error) int {
	return 0
}
