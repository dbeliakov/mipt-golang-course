package ini

import (
	"errors"
)

var ErrFileIsMalformed = errors.New("file is malformed")

func Parse(fileName string) (map[string]map[string]string, error) {
    panic("implement me")
}
