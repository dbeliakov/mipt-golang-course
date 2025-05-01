package mockhelper_test

import (
	"context"
	"errors"
	"testing"

	"github.com/dbeliakov/mipt-golang-course/tasks/06/mockhelper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------

type DB interface {
	GetUserIDs(context.Context) ([]int, error)
	DeleteUserByID(context.Context, int) error
}

var _ DB = &mockDB{}

type mockDB struct {
	mh mockhelper.MockHelper
}

func NewMockDB(mh mockhelper.MockHelper) *mockDB {
	return &mockDB{
		mh: mh,
	}
}

func (m *mockDB) GetUserIDs(ctx context.Context) ([]int, error) {
	returnValues := m.mh.Call("GetUserIDs", ctx)

	var ret0 []int
	if returnValues[0] != nil {
		ret0 = returnValues[0].([]int)
	}
	var ret1 error
	if returnValues[1] != nil {
		ret1 = returnValues[1].(error)
	}

	return ret0, ret1
}

func (m *mockDB) DeleteUserByID(ctx context.Context, id int) error {
	retVal := m.mh.Call("DeleteUserByID", ctx, id)

	var ret error
	if retVal[0] != nil {
		ret = retVal[0].(error)
	}
	return ret
}

// ----------------------------------

func DeleteAllUsers(ctx context.Context, db DB) []error {
	users, err := db.GetUserIDs(ctx)
	if err != nil {
		return []error{err}
	}

	var errs []error
	for _, id := range users {
		err = db.DeleteUserByID(ctx, id)
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

func TestDeleteAllUsers(t *testing.T) {
	ctx := context.Background()
	mh := mockhelper.NewMockHelper(t)
	db := NewMockDB(mh)

	expectIDs := []int{1, 2, 3}
	mh.ExpectCall("GetUserIDs", ctx).Return(expectIDs, nil)
	for _, id := range expectIDs {
		mh.ExpectCall("DeleteUserByID", ctx, id).Return(nil)
	}

	errs := DeleteAllUsers(ctx, db)
	assert.Empty(t, errs)
}

func TestDeleteAllUsers_GetIDsError(t *testing.T) {
	ctx := context.Background()
	mh := mockhelper.NewMockHelper(t)
	db := NewMockDB(mh)

	expectErr := errors.New("oops")
	mh.ExpectCall("GetUserIDs", ctx).Return(nil, expectErr)

	errs := DeleteAllUsers(ctx, db)
	require.Len(t, errs, 1)
	assert.Equal(t, expectErr, errs[0])
}

func TestDeleteAllUsers_DeleteErrors(t *testing.T) {
	ctx := context.Background()
	mh := mockhelper.NewMockHelper(t)
	db := NewMockDB(mh)

	expectIDs := []int{1, 2, 3}
	mh.ExpectCall("GetUserIDs", ctx).Return(expectIDs, nil)
	expectErr := errors.New("something went wrong")
	for _, id := range expectIDs {
		if id == 2 {
			mh.ExpectCall("DeleteUserByID", ctx, id).Return(nil)
			continue
		}
		mh.ExpectCall("DeleteUserByID", ctx, id).Return(expectErr)
	}

	errs := DeleteAllUsers(ctx, db)
	require.Len(t, errs, 2)
	assert.Equal(t, expectErr, errs[0])
	assert.Equal(t, expectErr, errs[1])
}

func TestMockHelper_AnyWorks(t *testing.T) {
	ctx := context.Background()
	mh := mockhelper.NewMockHelper(t)
	db := NewMockDB(mh)

	for i := 0; i < 10; i++ {
		mh.ExpectCall("DeleteUserByID", ctx, mockhelper.Any()).Return(nil)
	}

	for i := 0; i < 10; i++ {
		_ = db.DeleteUserByID(ctx, i)
	}
}
