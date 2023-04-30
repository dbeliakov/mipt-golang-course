package sign

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

type HMACTestCase struct {
	Content   []byte
	Secret    []byte
	Signature string
}

var testCasesHMAC = []HMACTestCase{
	{
		Content:   []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`),
		Secret:    []byte("dnfi13488fb20ncj"),
		Signature: "QNhtdE5wZLn35JvwAGbeXMfDxdK/2jq28Y0PBz5Bcfk=",
	},
	{
		Content:   []byte(`Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.`),
		Secret:    []byte("28n2idncas01"),
		Signature: "WgW3vMekaBv6T8ruzZp6BgghVeBzaTSacqdRTPN1VfU=",
	},
	{
		Content:   []byte(`But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness.`),
		Secret:    []byte("0924runfwnu9"),
		Signature: "kUA/tNuzcgRrwdbuncZVyEgYtUUyttBbLeRHdhqGFRw=",
	},
}

func TestHMACSHA256Signer(t *testing.T) {
	for i, tc := range testCasesHMAC {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			signer := NewHMACSigner(sha256.New, tc.Secret)
			sd, err := signer.Sign(tc.Content)
			require.NoError(t, err)
			require.Equal(t, Method(MethodHMAC), sd.Method)
			require.Equal(t, tc.Signature, sd.Signature)

			validator := NewHMACValidator(sha256.New, tc.Secret)
			ok, err := validator.Validate(sd)
			require.NoError(t, err)
			require.True(t, ok)

			sd.Method = MethodRSA
			_, err = validator.Validate(sd)
			require.ErrorIs(t, err, ErrInvalidMethod)
			sd.Method = MethodHMAC

			sd.Content = append(sd.Content, []byte("42")...)
			ok, err = validator.Validate(sd)
			require.NoError(t, err)
			require.False(t, ok)
		})
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
