package rw

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountReader(t *testing.T) {
	cr := NewCountReader()
	buf := make([]byte, 10)

	n, err := cr.Read(buf)
	require.NoError(t, err)
	assert.Equal(t, 10, n)
	expected := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, buf)

	n, err = cr.Read(buf[:3])
	require.NoError(t, err)
	assert.Equal(t, 3, n)
	assert.Equal(t, []byte{0, 1, 2}, buf[:3])

	n, err = cr.Read(buf[:2])
	require.NoError(t, err)
	assert.Equal(t, 2, n)
	assert.Equal(t, []byte{3, 4}, buf[:2])
}

func TestLimitReader(t *testing.T) {
	sr := strings.NewReader("Hello, World!")
	lr := NewLimitReader(sr, 5)

	data, err := io.ReadAll(lr)
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrLimitExceeded)
	assert.Equal(t, "Hello", string(data))

	sr = strings.NewReader("short")
	lr = NewLimitReader(sr, 6)

	data, err = io.ReadAll(lr)
	require.NoError(t, err)
	assert.Equal(t, "short", string(data))
}

func TestConcatReader(t *testing.T) {
	sr1 := strings.NewReader("Hello, ")
	sr2 := strings.NewReader("World!")
	cr := NewConcatReader(sr1, sr2)

	data, err := io.ReadAll(cr)
	require.NoError(t, err)
	assert.Equal(t, "Hello, World!", string(data))
}

func TestHexWriter(t *testing.T) {
	var buf bytes.Buffer
	hw := NewHexWriter(&buf)

	n, err := hw.Write([]byte("Hello"))
	require.NoError(t, err)
	assert.Equal(t, 10, n)
	assert.Equal(t, "48656c6c6f", buf.String())
}

func TestTeeWriter(t *testing.T) {
	var buf1, buf2 bytes.Buffer
	tw := NewTeeWriter(&buf1, &buf2)

	n, err := tw.Write([]byte("Hello"))
	require.NoError(t, err)
	assert.Equal(t, 5, n)
	assert.Equal(t, "Hello", buf1.String())
	assert.Equal(t, "Hello", buf2.String())
}
