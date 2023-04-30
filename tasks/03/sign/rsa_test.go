package sign

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
)

type RSATestCase struct {
	Content []byte
	Seed    int64
}

var testCasesRSA = []RSATestCase{
	{
		Content: []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`),
		Seed:    42,
	},
	{
		Content: []byte(`Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo.`),
		Seed:    192,
	},
	{
		Content: []byte(`But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the system, and expound the actual teachings of the great explorer of the truth, the master-builder of human happiness.`),
		Seed:    3820891,
	},
}

func TestRSASigner(t *testing.T) {
	for i, tc := range testCasesRSA {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			privateKey, err := rsa.GenerateKey(rand.New(rand.NewSource(tc.Seed)), 512)
			require.NoError(t, err)

			signer := NewRSASigner(crypto.SHA256, sha256.New, privateKey)
			sd, err := signer.Sign(tc.Content)
			require.NoError(t, err)
			require.Equal(t, Method(MethodRSA), sd.Method)

			validator := NewRSAValidator(crypto.SHA256, sha256.New, &privateKey.PublicKey)
			ok, err := validator.Validate(sd)
			require.NoError(t, err)
			require.True(t, ok)

			sd.Method = MethodHMAC
			_, err = validator.Validate(sd)
			require.ErrorIs(t, err, ErrInvalidMethod)
			sd.Method = MethodRSA

			sd.Content = append(sd.Content, []byte("42")...)
			ok, err = validator.Validate(sd)
			require.NoError(t, err)
			require.False(t, ok)
		})
	}
}
