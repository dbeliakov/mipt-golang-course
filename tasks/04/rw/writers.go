package rw

import "io"

type HexWriter struct{}

func NewHexWriter(w io.Writer) *HexWriter {
	panic("implement me")
}

func (h *HexWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}

type TeeWriter struct{}

func NewTeeWriter(writers ...io.Writer) *TeeWriter {
	panic("implement me")
}

func (t *TeeWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}
