package rw

import (
	"errors"
	"io"
)

type CountReader struct {
}

func NewCountReader() *CountReader {
	return &CountReader{}
}

var ErrLimitExceeded = errors.New("limit exceeded")

type LimitReader struct {
}

func NewLimitReader(r io.Reader, limit int) *LimitReader {
	return &LimitReader{}
}

type ConcatReader struct {
}

func NewConcatReader(rs ...io.Reader) *ConcatReader {
	return &ConcatReader{}
}
