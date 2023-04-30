package sign

import "errors"

type Method string

const (
	MethodHMAC = "HMAC"
	MethodRSA  = "RSA"
)

type SignedData struct {
	Method    Method `json:"method"`
	Signature string `json:"signature"`
	Content   []byte `json:"content"`
}

type Signer interface {
	Sign(data []byte) (SignedData, error)
}

var (
	ErrInvalidMethod = errors.New("invalid signing method")
)

type Validator interface {
	Validate(data SignedData) (bool, error)
}
