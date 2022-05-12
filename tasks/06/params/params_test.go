package params

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	var s testStruct
	var v = url.Values{}
	v.Set("string", "hello")
	v.Set("int-value", "42")
	v.Set("bool1", "true")
	v.Set("bool2", "true")
	v.Set("slice-of-ints", "1")
	v.Add("slice-of-ints", "2")
	v.Add("slice-of-ints", "3")
	v.Add("private", "42")
	v.Add("privateunsupported", "unexpected")
	v.Add("ptr-string", "world")
	v.Add("i", "1")

	require.ErrorIs(t, ErrNotPointer, Unpack(v, s))
	require.NoError(t, Unpack(v, &s))

	require.Equal(t, "hello", s.String)
	require.Equal(t, 42, s.Int)
	require.True(t, s.Bool1)
	require.True(t, s.Bool2)
	require.Equal(t, []int{1, 2, 3}, s.Slice)
	require.Equal(t, 0, s.private)
	require.NotNil(t, s.Ptr)
	require.Equal(t, "world", *s.Ptr)

	require.Error(t, Unpack(v, &testStructUnsupported{}))
}

type testStruct struct {
	String             string `http:"string"`
	Int                int    `http:"int-value"`
	Bool1              bool   `http:"bool1"`
	Bool2              bool
	Slice              []int    `http:"slice-of-ints"`
	private            int      `http:"private"`
	privateUnsupported struct{} `http:"privateunsupported"`
	Ptr                *string  `http:"ptr-string"`
}

type testStructUnsupported struct {
	I struct{}
}
