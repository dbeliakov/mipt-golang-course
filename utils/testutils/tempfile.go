package testutils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TempFile(t *testing.T, content string) (fileName string) {
	t.Helper()

	f, err := os.CreateTemp(t.TempDir(), "test-*")
	require.NoError(t, err)
	_, err = f.WriteString(content)
	require.NoError(t, err)

	require.NoError(t, f.Close())
	t.Cleanup(func() {
		_ = os.Remove(f.Name())
	})

	return f.Name()
}
