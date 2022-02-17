package hello_test

import (
	"github.com/stretchr/testify/require"

	"github.com/dbeliakov/mipt-golang-course/tasks/01/hello"

	"testing"
)

func TestHello(t *testing.T) {
	require.Equal(t, "Hello, world!", hello.Hello())
}