package polish

import (
	"errors"
)

var ErrInvalidExpression = errors.New("invalid expression")

func Calculate(expr string) (int, error) {
	panic("implement me")
}
