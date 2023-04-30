package sign

import (
	"hash"
)

type HMACSHA256Signer struct {
}

func NewHMACSigner(hashF func() hash.Hash, secret []byte) *HMACSHA256Signer {
	return &HMACSHA256Signer{}
}

var _ Signer = &HMACSHA256Signer{}

func (s *HMACSHA256Signer) Sign(data []byte) (SignedData, error) {
	return SignedData{}, nil
}

type HMACSHA256Validator struct {
}

var _ Validator = &HMACSHA256Validator{}

func NewHMACValidator(hashF func() hash.Hash, secret []byte) *HMACSHA256Validator {
	return &HMACSHA256Validator{}
}

func (s *HMACSHA256Validator) Validate(data SignedData) (bool, error) {
	return true, nil
}
