package rw

import (
	"errors"
	"io"
)

type CountReader struct{}

func NewCountReader() *CountReader {
	return &CountReader{}
}

func (c *CountReader) Read(p []byte) (int, error) {
	panic("implement me")
}

var ErrLimitExceeded = errors.New("limit exceeded")

type LimitReader struct{}

func NewLimitReader(r io.Reader, limit int) *LimitReader {
	panic("implement me")
}

func (l *LimitReader) Read(p []byte) (int, error) {
	panic("implement me")
}

type ConcatReader struct{}

func NewConcatReader(rs ...io.Reader) *ConcatReader {
	panic("implement me")
}

// Read склеивает чтение из цепочки io.Reader по аналогии с io.MultiReader:
// при EOF у текущего читателя переходит к следующему; при неполном Read заполняет буфер дальше.
func (c *ConcatReader) Read(p []byte) (n int, err error) {
	panic("implement me")
}
