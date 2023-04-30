package sign

import (
	"crypto"
	"crypto/rsa"
	"hash"
)

type RSASigner struct {
}

func NewRSASigner(hash crypto.Hash, hashF func() hash.Hash, privateKey *rsa.PrivateKey) *RSASigner {
	return &RSASigner{}
}

var _ Signer = &RSASigner{}

func (s RSASigner) Sign(data []byte) (SignedData, error) {
	return SignedData{}, nil
}

type RSAValidator struct {
}

func NewRSAValidator(hash crypto.Hash, hashF func() hash.Hash, publicKey *rsa.PublicKey) *RSAValidator {
	return &RSAValidator{}
}

var _ Validator = &RSAValidator{}

func (v *RSAValidator) Validate(data SignedData) (bool, error) {
	return true, nil
}
