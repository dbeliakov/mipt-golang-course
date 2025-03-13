package jsoncensor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCensorJSON(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		substr       string
		expectedJSON string
		expectError  bool
	}{
		{
			name:         "simple key-value pair",
			input:        []byte(`{"k": "v", "key": "word_to_censor"}`),
			substr:       "word_to_censor",
			expectedJSON: `{"k": "v", "key": "***"}`,
			expectError:  false,
		},
		{
			name:         "nested JSON object",
			input:        []byte(`{"k": "v", "another_key": {"arr": ["word_to_censor"], "inner_key": "what if i have a word_to_censor here???"}}`),
			substr:       "word_to_censor",
			expectedJSON: `{"k": "v", "another_key": {"arr": ["***"], "inner_key": "***"}}`,
			expectError:  false,
		},
		{
			name:         "array of strings",
			input:        []byte(`["word_to_censor", "safe_word", "another_word_to_censor"]`),
			substr:       "word_to_censor",
			expectedJSON: `["***", "safe_word", "***"]`,
			expectError:  false,
		},
		{
			name:         "nested arrays",
			input:        []byte(`{"arr": [["word_to_censor", "safe_word"], ["another_word_to_censor"]]}`),
			substr:       "word_to_censor",
			expectedJSON: `{"arr": [["***", "safe_word"], ["***"]]}`,
			expectError:  false,
		},
		{
			name:         "no match",
			input:        []byte(`{"k": "v", "key": "safe_value"}`),
			substr:       "word_to_censor",
			expectedJSON: `{"k": "v", "key": "safe_value"}`,
			expectError:  false,
		},
		{
			name:         "invalid JSON",
			input:        []byte(`{"k": "v", "key": "word_to_censor"`), // Missing closing brace
			substr:       "word_to_censor",
			expectedJSON: "",
			expectError:  true,
		},
		{
			name:         "empty JSON",
			input:        []byte(`{}`),
			substr:       "word_to_censor",
			expectedJSON: `{}`,
			expectError:  false,
		},
		{
			name:         "empty string",
			input:        []byte(``),
			substr:       "word_to_censor",
			expectedJSON: "",
			expectError:  true,
		},
		{
			name:         "complex nested JSON",
			input:        []byte(`{"k1": "v1", "k2": {"k3": "word_to_censor", "k4": [{"k5": "word_to_censor"}, {"k6": "safe_value"}]}, "k7": ["word_to_censor", "safe_word"]}`),
			substr:       "word_to_censor",
			expectedJSON: `{"k1": "v1", "k2": {"k3": "***", "k4": [{"k5": "***"}, {"k6": "safe_value"}]}, "k7": ["***", "safe_word"]}`,
			expectError:  false,
		},
		{
			name:         "deeply nested JSON",
			input:        []byte(`{"k1": {"k2": {"k3": {"k4": "word_to_censor"}}}}`),
			substr:       "word_to_censor",
			expectedJSON: `{"k1": {"k2": {"k3": {"k4": "***"}}}}`,
			expectError:  false,
		},
		{
			name:         "mixed types in array",
			input:        []byte(`{"arr": ["word_to_censor", 123, true, {"k": "word_to_censor"}]}`),
			substr:       "word_to_censor",
			expectedJSON: `{"arr": ["***", 123, true, {"k": "***"}]}`,
			expectError:  false,
		},
		{
			name:         "multiple occurrences in one string",
			input:        []byte(`{"k": "word_to_censor and another word_to_censor"}`),
			substr:       "word_to_censor",
			expectedJSON: `{"k": "***"}`,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := CensorJSON(tt.input, tt.substr)

			if tt.expectError {
				assert.Error(t, err, "Expected an error")
				return
			}

			require.NoError(t, err)
			assert.JSONEq(t, tt.expectedJSON, string(res))
		})
	}
}
