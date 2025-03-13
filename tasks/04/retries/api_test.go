package retries

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockAPI struct {
	values          map[string]*Value
	epochs          map[string]uint64
	simulateErrors  sync.Map
	writeErrors     sync.Map
	simulateSetFail bool
}

func NewMockAPI() *MockAPI {
	return &MockAPI{
		values: make(map[string]*Value),
		epochs: make(map[string]uint64),
	}
}

func (m *MockAPI) Get(key string) (*Value, uint64, error) {
	err, ok := m.simulateErrors.Load(key)
	if ok {
		return nil, 0, err.(error)
	}

	val, ok := m.values[key]
	if !ok {
		return nil, 0, APIError{status: StatusNotFound}
	}
	return val, m.epochs[key], nil
}

func (m *MockAPI) Set(key string, targetEpoch uint64, value *Value) error {
	err, ok := m.writeErrors.Load(key)
	if ok {
		return err.(error)
	}
	if m.simulateSetFail {
		return APIError{status: StatusFatalError}
	}
	if targetEpoch <= m.epochs[key] {
		return APIError{status: StatusValueTooOld}
	}
	m.values[key] = value
	m.epochs[key] = targetEpoch
	return nil
}

func TestUpdateValue_Success(t *testing.T) {
	api := NewMockAPI()
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(api, "key", updateFn)
	require.NoError(t, err, "Update should succeed")
	assert.Equal(t, uint64(2), api.epochs["key"], "Epoch should be incremented")
}

func TestUpdateValue_TransientErrors(t *testing.T) {
	api := NewMockAPI()
	api.simulateErrors.Store("key", ErrNetworkFault)
	api.values["key"] = &Value{}

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	// Simulate a successful update after a few retries
	go func() {
		time.Sleep(100 * time.Millisecond)
		api.simulateErrors.Delete("key")
	}()

	err := UpdateValue(api, "key", updateFn)
	require.NoError(t, err, "Update should succeed after retries")
}

func TestUpdateValue_ValueTooOld(t *testing.T) {
	api := NewMockAPI()
	api.writeErrors.Store("key", APIError{status: StatusValueTooOld})
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	// Simulate a concurrent update
	go func() {
		time.Sleep(100 * time.Millisecond)
		api.writeErrors.Delete("key")
	}()

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(api, "key", updateFn)
	require.NoError(t, err)
	assert.Equal(t, uint64(2), api.epochs["key"])
}

func TestUpdateValue_KeyNotFound(t *testing.T) {
	api := NewMockAPI()

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(api, "nonexistent", updateFn)
	require.Error(t, err)
	assert.ErrorAs(t, err, &APIError{})
	assert.ErrorIs(t, err, APIError{status: StatusNotFound})
}

func TestUpdateValue_FatalError(t *testing.T) {
	api := NewMockAPI()
	api.values["key"] = &Value{}
	api.epochs["key"] = 1
	api.simulateSetFail = true

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(api, "key", updateFn)
	require.Error(t, err)
	assert.ErrorAs(t, err, &APIError{})
	assert.ErrorIs(t, err, APIError{status: StatusFatalError})
}

func TestUpdateValue_ValueUpdaterFailure(t *testing.T) {
	api := NewMockAPI()
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	updateFn := func(currentValue *Value) (*Value, error) {
		return nil, errors.New("update failed")
	}

	err := UpdateValue(api, "key", updateFn)
	require.Error(t, err)
}
