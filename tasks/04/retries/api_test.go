package retries

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockAPI struct {
	mu     sync.Mutex
	values map[string]*Value
	epochs map[string]uint64
	simulateErrors sync.Map
	writeErrors    sync.Map
	simulateSetFail bool

	transientGetFails int
	transientGetErr   error

	transientSetValueTooOld int
}

func NewMockAPI() *MockAPI {
	return &MockAPI{
		values: make(map[string]*Value),
		epochs: make(map[string]uint64),
	}
}

func (m *MockAPI) Get(key string) (*Value, uint64, error) {
	m.mu.Lock()
	if m.transientGetFails > 0 {
		m.transientGetFails--
		err := m.transientGetErr
		if err == nil {
			err = ErrNetworkFault
		}
		m.mu.Unlock()
		return nil, 0, err
	}
	m.mu.Unlock()

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
	m.mu.Lock()
	if m.transientSetValueTooOld > 0 {
		m.transientSetValueTooOld--
		m.mu.Unlock()
		return APIError{status: StatusValueTooOld}
	}
	m.mu.Unlock()

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

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{})
	require.NoError(t, err, "Update should succeed")
	assert.Equal(t, uint64(2), api.epochs["key"], "Epoch should be incremented")
}

func TestUpdateValue_TransientErrors(t *testing.T) {
	api := NewMockAPI()
	api.transientGetFails = 5
	api.values["key"] = &Value{}

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{MaxRetries: defaultMaxRetries})
	require.NoError(t, err, "Update should succeed after retries")
}

func TestUpdateValue_ValueTooOld(t *testing.T) {
	api := NewMockAPI()
	api.transientSetValueTooOld = 1
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{MaxRetries: defaultMaxRetries})
	require.NoError(t, err)
	assert.Equal(t, uint64(2), api.epochs["key"])
}

func TestUpdateValue_KeyNotFound(t *testing.T) {
	api := NewMockAPI()

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(context.Background(), api, "nonexistent", updateFn, UpdateOptions{})
	require.Error(t, err)
	var apiErr APIError
	assert.True(t, errors.As(err, &apiErr))
	assert.Equal(t, StatusNotFound, apiErr.Status())
}

func TestUpdateValue_FatalError(t *testing.T) {
	api := NewMockAPI()
	api.values["key"] = &Value{}
	api.epochs["key"] = 1
	api.simulateSetFail = true

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{})
	require.Error(t, err)
	var apiErr APIError
	assert.True(t, errors.As(err, &apiErr))
	assert.Equal(t, StatusFatalError, apiErr.Status())
}

func TestUpdateValue_ValueUpdaterFailure(t *testing.T) {
	api := NewMockAPI()
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	updateFn := func(currentValue *Value) (*Value, error) {
		return nil, errors.New("update failed")
	}

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{})
	require.Error(t, err)
}

func TestUpdateValue_RetryBudgetExhausted(t *testing.T) {
	api := NewMockAPI()
	api.simulateErrors.Store("key", ErrNetworkFault)
	api.values["key"] = &Value{}

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{MaxRetries: 3})
	require.Error(t, err)
	assert.ErrorIs(t, err, ErrRetryBudgetExhausted)
}

func TestUpdateValue_ContextCancelled(t *testing.T) {
	api := NewMockAPI()
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(ctx, api, "key", updateFn, UpdateOptions{})
	require.Error(t, err)
	assert.ErrorIs(t, err, context.Canceled)
}

func TestUpdateValue_OnRetry(t *testing.T) {
	api := NewMockAPI()
	api.transientGetFails = 1
	api.transientGetErr = ErrTimeout
	api.values["key"] = &Value{}
	api.epochs["key"] = 1

	var attempts []int
	var errs []error
	updateFn := func(currentValue *Value) (*Value, error) {
		return &Value{}, nil
	}

	err := UpdateValue(context.Background(), api, "key", updateFn, UpdateOptions{
		MaxRetries: defaultMaxRetries,
		OnRetry: func(attempt int, e error) {
			attempts = append(attempts, attempt)
			errs = append(errs, e)
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, attempts)
	assert.Equal(t, 1, attempts[0])
	assert.ErrorIs(t, errs[0], ErrTimeout)
}
