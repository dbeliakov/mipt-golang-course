package rw

import (
	"io"
)

type HexWriter struct {
}

func NewHexWriter(w io.Writer) *HexWriter {
	return &HexWriter{}
}

type TeeWriter struct {
}

func NewTeeWriter(writers ...io.Writer) *TeeWriter {
	return &TeeWriter{}
}
