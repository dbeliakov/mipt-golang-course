package ini

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dbeliakov/mipt-golang-course/utils/testutils"
)

func TestParse_ErrorOnNonExistentFile(t *testing.T) {
	// NOTE: file with such name cannot exist, because file name is too long
	nonExistentFile := strings.Repeat("abcd", 512)

	_, err := Parse(nonExistentFile)
	require.Error(t, err, "Parser should return an error when file does not exist")
}

func TestParse_CommonCases(t *testing.T) {
	testCases := []struct {
		name          string
		content       string
		expected      map[string]map[string]string
		invalidFormat bool
	}{
		{
			name:     "empty",
			content:  "",
			expected: map[string]map[string]string{},
		},
		{
			name:     "single header without keys",
			content:  "[header1]",
			expected: map[string]map[string]string{},
		},
		{
			name: "single header with keys",
			content: `
[some_header]
key1 = value
k2=   val2
			`,
			expected: map[string]map[string]string{
				"some_header": {
					"key1": "value",
					"k2":   "val2",
				},
			},
		},
		{
			name: "empty lines inside header",
			content: `
[header]
k1 = v1




k2 = v2
			`,
			expected: map[string]map[string]string{
				"header": {
					"k1": "v1",
					"k2": "v2",
				},
			},
		},
		{
			name: "multiple headers without keys",
			content: `
[header1]
[header2]
[header3]
			`,
			expected: map[string]map[string]string{},
		},
		{
			name: "multiple headers with keys",
			content: `
[1]
key = first_value

[second]
key = 2_value

[third]
key = third_value

[4]
key = 4
			`,
			expected: map[string]map[string]string{
				"1": {
					"key": "first_value",
				},
				"second": {
					"key": "2_value",
				},
				"third": {
					"key": "third_value",
				},
				"4": {
					"key": "4",
				},
			},
		},
		{
			name: "duplicate keys are overwritten in single header",
			content: `
[header]
duplicate_key = first_value
unique_key = unique_value
duplicate_key = last_value
			`,
			expected: map[string]map[string]string{
				"header": {
					"duplicate_key": "last_value",
					"unique_key":    "unique_value",
				},
			},
		},
		{
			name: "duplicate headers are merged together",
			content: `
[duplicate_header]
key = value

[unique_header]
unique_key = some_unique_value

[duplicate_header]
another_key = another_value
			`,
			expected: map[string]map[string]string{
				"duplicate_header": {
					"key":         "value",
					"another_key": "another_value",
				},
				"unique_header": {
					"unique_key": "some_unique_value",
				},
			},
		},
		{
			name: "duplicate keys are overwritten in duplicate header when merged",
			content: `
[header]
k = 1

[header]
k = 2

[another_header]
k = 3

[header]
k = 4
			`,
			expected: map[string]map[string]string{
				"header": {
					"k": "4",
				},
				"another_header": {
					"k": "3",
				},
			},
		},
		{
			name: "error on empty header inside brackets",
			content: `
[]
value = 10
			`,
			invalidFormat: true,
		},
		{
			name: "error on multiple equal signs inside key-value definition",
			content: `
[header]
a=b=c
			`,
			invalidFormat: true,
		},
		{
			name: "error on key-value pair outside headers",
			content: `
key = value

[header]
another_key = value
			`,
			invalidFormat: true,
		},
		{
			name: "error on key without a value",
			content: `
[header]
key
			`,
			invalidFormat: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fileName := testutils.TempFile(t, tc.content)
			actual, err := Parse(fileName)
			if tc.invalidFormat {
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrFileIsMalformed)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
