package params

import (
	"errors"
	"net/url"
)

var (
	ErrNotPointer = errors.New("object not a pointer")
)

func Unpack(values url.Values, to interface{}) error {
	return nil
}
