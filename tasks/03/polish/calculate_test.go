package polish

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type WantResult int

const (
	Value                WantResult = 0
	AnyErr               WantResult = 1
	InvalidExpressionErr WantResult = 2
)

func TestCalculate(t *testing.T) {
	testCases := []struct {
		expr          string
		expectedValue int
		wantResult    WantResult
	}{
		{expr: "", expectedValue: 0},
		{expr: "1 2 +", expectedValue: 3},
		{expr: "12 15 -", expectedValue: -3},
		{expr: "1 4 56 + +", expectedValue: 61},
		{expr: "-3 4 12 -15 - * *", expectedValue: -324},
		{expr: "a b +", wantResult: AnyErr},
		{expr: "3 4 5 +", wantResult: InvalidExpressionErr},
		{expr: "1 *", wantResult: InvalidExpressionErr},
	}

	for _, tc := range testCases {
		t.Run(tc.expr, func(t *testing.T) {
			res, err := Calculate(tc.expr)

			switch tc.wantResult {
			case Value:
				require.NoError(t, err)
				assert.Equal(t, tc.expectedValue, res)
			case AnyErr:
				assert.Error(t, err)
			case InvalidExpressionErr:
				require.Error(t, err)
				assert.ErrorIs(t, err, ErrInvalidExpression)
			}
		})
	}
}
