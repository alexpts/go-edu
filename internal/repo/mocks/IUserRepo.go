// Code generated by mockery v2.33.3. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/alexpts/edu-go/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// IUserRepoMock is an autogenerated mock type for the IUserRepo type
type IUserRepoMock struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *IUserRepoMock) Create(ctx context.Context, _a1 *model.User) (*model.User, int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.User
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*model.User, int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) int64); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.User) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindAll provides a mock function with given fields: ctx, relations
func (_m *IUserRepoMock) FindAll(ctx context.Context, relations ...string) ([]model.User, error) {
	_va := make([]interface{}, len(relations))
	for _i := range relations {
		_va[_i] = relations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) ([]model.User, error)); ok {
		return rf(ctx, relations...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []model.User); ok {
		r0 = rf(ctx, relations...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, relations...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByNameRawSQL provides a mock function with given fields: name
func (_m *IUserRepoMock) FindByNameRawSQL(name string) (*model.User, error) {
	ret := _m.Called(name)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneById provides a mock function with given fields: ctx, id, relations
func (_m *IUserRepoMock) FindOneById(ctx context.Context, id int, relations ...string) (*model.User, error) {
	_va := make([]interface{}, len(relations))
	for _i := range relations {
		_va[_i] = relations[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, id)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, ...string) (*model.User, error)); ok {
		return rf(ctx, id, relations...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, ...string) *model.User); ok {
		r0 = rf(ctx, id, relations...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, ...string) error); ok {
		r1 = rf(ctx, id, relations...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindOneUserByName provides a mock function with given fields: name
func (_m *IUserRepoMock) FindOneUserByName(name string) (*model.User, error) {
	ret := _m.Called(name)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Persist provides a mock function with given fields: ctx, _a1
func (_m *IUserRepoMock) Persist(ctx context.Context, _a1 *model.User) (*model.User, int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.User
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*model.User, int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) int64); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.User) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *IUserRepoMock) Update(ctx context.Context, _a1 *model.User) (*model.User, int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *model.User
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) (*model.User, int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) *model.User); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) int64); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.User) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewIUserRepoMock creates a new instance of IUserRepoMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIUserRepoMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *IUserRepoMock {
	mock := &IUserRepoMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}