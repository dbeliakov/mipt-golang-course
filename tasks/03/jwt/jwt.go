package jwt

import (
	"errors"
	"time"
)

type SignMethod string

const (
	HS256 SignMethod = "HS256"
	HS512 SignMethod = "HS512"
)

var (
	ErrInvalidSignMethod      = errors.New("invalid sign method")
	ErrSignatureInvalid       = errors.New("signature invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrSignMethodMismatched   = errors.New("sign method mismatched")
	ErrConfigurationMalformed = errors.New("configuration malformed")
	ErrInvalidToken           = errors.New("invalid token")
)

func Encode(data interface{}, opts ...Option) ([]byte, error) {
	// TODO: implement me
	return nil, nil
}

func Decode(token []byte, data interface{}, opts ...Option) error {
	// TODO: implement me
	return nil
}

// To mock time in tests
var timeFunc = time.Now
