package pr

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dbeliakov/mipt-golang-course/utils/testutils"
)

const binPath = "github.com/dbeliakov/mipt-golang-course/tasks/03/pr/cmd/gopr"

func TestPR(t *testing.T) {
	testCases := []struct {
		file    string
		nLines  int
		wantErr bool
	}{
		{file: filepath.Join("01", "simple.txt"), nLines: 4},
		{file: filepath.Join("02", "multiple_pages.txt"), nLines: 4},
		{file: filepath.Join("03", "add_newlines.txt"), nLines: 4},
		// NOTE: this file does not exist
		{file: strings.Repeat("abcd", 128), wantErr: true},
	}

	binCache := testutils.NewBinaryCache()
	testBinary := binCache.LoadBinary(binPath)

	for _, tc := range testCases {
		t.Run(tc.file, func(t *testing.T) {
			inputFile := filepath.Join("testdata", tc.file)
			expectedOutputFile := filepath.Join("testdata", tc.file+".output")

			cmd := exec.Command(testBinary, "-n", strconv.Itoa(tc.nLines), inputFile)
			res, err := cmd.Output()
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				expectedResult, err := os.ReadFile(expectedOutputFile)
				require.NoError(t, err)

				if !assert.Equal(t, string(expectedResult), string(res)) {
					fileNameToDumpResult := inputFile + ".actual"
					err = os.WriteFile(fileNameToDumpResult, res, 0644)
					require.NoError(t, err)

					t.Logf("Actual result has been written to %q", fileNameToDumpResult)
				}
			}
		})
	}
}

func TestPR_MultipleFiles(t *testing.T) {
	args := []string{
		"-n", "6",
		filepath.Join("testdata", "multiple_files", "a.txt"),
		filepath.Join("testdata", "multiple_files", "b.txt"),
		filepath.Join("testdata", "multiple_files", "c.txt"),
	}
	expectedOutputFile := filepath.Join("testdata", "multiple_files", "expected.txt")

	binCache := testutils.NewBinaryCache()
	testBinary := binCache.LoadBinary(binPath)

	cmd := exec.Command(testBinary, args...)
	res, err := cmd.Output()
	require.NoError(t, err)

	expectedResult, err := os.ReadFile(expectedOutputFile)
	require.NoError(t, err)

	if !assert.Equal(t, string(expectedResult), string(res)) {
		fileNameToDumpResult := filepath.Join("testdata", "multiple_files", "actual.txt")
		err = os.WriteFile(fileNameToDumpResult, res, 0644)
		require.NoError(t, err)

		t.Logf("Actual result has been written to %q", fileNameToDumpResult)
	}
}
