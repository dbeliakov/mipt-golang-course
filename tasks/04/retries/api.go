package retries

import (
	"errors"
)

var (
	ErrTimeout      = errors.New("timeout")
	ErrNetworkFault = errors.New("network error")
)

type Status string

const (
	StatusNotFound    Status = "not found"
	StatusValueTooOld Status = "value too old"
	StatusFatalError  Status = "fatal error"
)

type APIError struct {
	status Status
}

func (a APIError) Error() string {
	return "encountered api error: " + string(a.status)
}

func (a APIError) Status() Status {
	return a.status
}

type Value struct{}

type SimpleAPI interface {
	Get(key string) (val *Value, epoch uint64, err error)
	Set(key string, targetEpoch uint64, value *Value) error
}

type ValueUpdater func(currentValue *Value) (*Value, error)

func UpdateValue(api SimpleAPI, key string, update ValueUpdater) error {
	panic("implement me")
}
