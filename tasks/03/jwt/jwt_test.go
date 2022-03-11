package jwt

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var EncodeTestCases = []struct {
	Data  interface{}
	Opts  []Option
	Token string
	Err   error
	Now   *time.Time
}{
	{
		Data: map[string]interface{}{
			"name":  "Dmitrii",
			"score": 42,
		},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("secret-key1"))},
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkIjp7Im5hbWUiOiJEbWl0cmlpIiwic2NvcmUiOjQyfX0.i1bDgy6dXTkMj13wZmJ5-dO4Jwq8oY9qQpUTHCTwR7Q",
	},
	{
		Data:  "hello, world",
		Opts:  []Option{WithSignMethod(HS512), WithKey([]byte("secret-key2"))},
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJkIjoiaGVsbG8sIHdvcmxkIn0.dVKVL5lqL_Wt7p1NSkk0r4CAyckSDuserSKNB3OkgTfvNsTcKsZ0MnD-rE38zflsYSWaXok4xwLeYjK6nWgcyQ",
	},
	{
		Data: struct {
			UserID string `json:"user-id"`
		}{
			UserID: "0x224",
		},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("secret-key3")), WithTTL(90 * time.Second)},
		Now:   timePtr(time.Unix(10, 0)),
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkIjp7InVzZXItaWQiOiIweDIyNCJ9LCJleHAiOjEwMH0.Asv19dLgMlNLIsQ-PFg3gjfIjFgAYAmZ4NFoUdYSHUk",
	},
	{
		Data:  42,
		Opts:  []Option{WithSignMethod(HS512), WithKey([]byte("secret-key4")), WithExpires(time.Unix(10, 0))},
		Now:   timePtr(time.Unix(5, 0)),
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJkIjo0MiwiZXhwIjoxMH0.3pgMRsfuTXl7xsXOtHJLCCsHVRJwrAEIvMuE4igFzAQDzOHJulSzpUYV0OnJDAXdHJ_a4PjvjcZpPp3pZ1kWzw",
	},
	{
		Opts: []Option{WithSignMethod("invalid")},
		Err:  ErrInvalidSignMethod,
	},
	{
		Opts: []Option{WithSignMethod(HS256), WithExpires(time.Now()), WithTTL(time.Second)},
		Err:  ErrConfigurationMalformed,
	},
	{
		Data: 42,
		Opts: []Option{WithSignMethod(HS512), WithKey([]byte("secret-key5")), WithExpires(time.Unix(5, 0))},
		Now:  timePtr(time.Unix(10, 0)),
		Err:  ErrConfigurationMalformed,
	},
}

func TestEncode(t *testing.T) {
	for i, tc := range EncodeTestCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tc.Now != nil {
				timeFunc = func() time.Time {
					return *tc.Now
				}
			} else {
				timeFunc = time.Now
			}
			token, err := Encode(tc.Data, tc.Opts...)
			if tc.Err != nil {
				require.ErrorIs(t, err, tc.Err)
			} else {
				require.NoError(t, err)
			}
			if tc.Token != "" {
				require.Equal(t, tc.Token, string(token))
			}
		})
	}
}

var DecodeTestCases = []struct {
	Data  interface{}
	Opts  []Option
	Token string
	Err   error
	Now   *time.Time
}{
	{
		Data: map[string]interface{}{
			"name":  "Dmitrii",
			"score": float64(42),
		},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("secret-key1"))},
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkIjp7Im5hbWUiOiJEbWl0cmlpIiwic2NvcmUiOjQyfX0.i1bDgy6dXTkMj13wZmJ5-dO4Jwq8oY9qQpUTHCTwR7Q",
	},
	{
		Data: map[string]interface{}{
			"name":  "Ivan",
			"score": float64(100),
		},
		Opts:  []Option{WithSignMethod(HS512), WithKey([]byte("secret-key2"))},
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJkIjp7Im5hbWUiOiJJdmFuIiwic2NvcmUiOjEwMH19.1JkKbC3SPeiwziZRcTWZDZN0tZ-uz2AKjiyUZsFm3xPlv9xifnKB8nckD5gVdG0ktkHG5P5nbACrhZth5u2ZPQ",
	},
	{
		Data:  map[string]interface{}{},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("1"))},
		Token: "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJkIjp7Im5hbWUiOiJJdmFuIiwic2NvcmUiOjEwMH19.1JkKbC3SPeiwziZRcTWZDZN0tZ-uz2AKjiyUZsFm3xPlv9xifnKB8nckD5gVdG0ktkHG5P5nbACrhZth5u2ZPQ",
		Err:   ErrSignMethodMismatched,
	},
	{
		Data:  map[string]interface{}{},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("1"))},
		Token: "123.eyJkIjp7Im5hbWUiOiJJdmFuIiwic2NvcmUiOjEwMH19.1JkKbC3SPeiwziZRcTWZDZN0tZ-uz2AKjiyUZsFm3xPlv9xifnKB8nckD5gVdG0ktkHG5P5nbACrhZth5u2ZPQ",
		Err:   ErrInvalidToken,
	},
	{
		Data:  map[string]interface{}{},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("secret-key"))},
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkIjp7Im5hbWUiOiJJdmFuIiwic2NvcmUiOjEwMH0sImV4cCI6MTB9.tQbJ0kB5pmbmXtPREVkVQupu5O3QlAXug1iJqnX2i44",
		Err:   ErrSignatureInvalid,
	},
	{
		Data:  map[string]interface{}{},
		Opts:  []Option{WithSignMethod(HS256), WithKey([]byte("secret-key"))},
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkIjp7Im5hbWUiOiJJdmFuIiwic2NvcmUiOjEwMH0sImV4cCI6MTB9.W0oIp8FzcNGB2HIZ5b_l6CrQSnhBpzKSbRrq-_JS9UM",
		Err:   ErrTokenExpired,
	},
}

func TestDecode(t *testing.T) {
	for i, tc := range DecodeTestCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if tc.Now != nil {
				timeFunc = func() time.Time {
					return *tc.Now
				}
			} else {
				timeFunc = time.Now
			}
			var data map[string]interface{}
			err := Decode([]byte(tc.Token), &data, tc.Opts...)
			if tc.Err != nil {
				require.ErrorIs(t, err, tc.Err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.Data, data)
			}
		})
	}
}

func timePtr(t time.Time) *time.Time {
	return &t
}
