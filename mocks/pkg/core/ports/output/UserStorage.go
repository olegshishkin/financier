// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/olegshishkin/financier/pkg/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserStorage is an autogenerated mock type for the UserStorage type
type UserStorage struct {
	mock.Mock
}

type UserStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *UserStorage) EXPECT() *UserStorage_Expecter {
	return &UserStorage_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: user
func (_m *UserStorage) Create(user *domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserStorage_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserStorage_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - user *domain.User
func (_e *UserStorage_Expecter) Create(user interface{}) *UserStorage_Create_Call {
	return &UserStorage_Create_Call{Call: _e.mock.On("Create", user)}
}

func (_c *UserStorage_Create_Call) Run(run func(user *domain.User)) *UserStorage_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.User))
	})
	return _c
}

func (_c *UserStorage_Create_Call) Return(_a0 error) *UserStorage_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindActiveByEmail provides a mock function with given fields: email
func (_m *UserStorage) FindActiveByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStorage_FindActiveByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindActiveByEmail'
type UserStorage_FindActiveByEmail_Call struct {
	*mock.Call
}

// FindActiveByEmail is a helper method to define mock.On call
//   - email string
func (_e *UserStorage_Expecter) FindActiveByEmail(email interface{}) *UserStorage_FindActiveByEmail_Call {
	return &UserStorage_FindActiveByEmail_Call{Call: _e.mock.On("FindActiveByEmail", email)}
}

func (_c *UserStorage_FindActiveByEmail_Call) Run(run func(email string)) *UserStorage_FindActiveByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserStorage_FindActiveByEmail_Call) Return(_a0 *domain.User, _a1 error) *UserStorage_FindActiveByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *UserStorage) Get(id string) (*domain.User, error) {
	ret := _m.Called(id)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStorage_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type UserStorage_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id string
func (_e *UserStorage_Expecter) Get(id interface{}) *UserStorage_Get_Call {
	return &UserStorage_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *UserStorage_Get_Call) Run(run func(id string)) *UserStorage_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserStorage_Get_Call) Return(_a0 *domain.User, _a1 error) *UserStorage_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Update provides a mock function with given fields: user
func (_m *UserStorage) Update(user *domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserStorage_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type UserStorage_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - user *domain.User
func (_e *UserStorage_Expecter) Update(user interface{}) *UserStorage_Update_Call {
	return &UserStorage_Update_Call{Call: _e.mock.On("Update", user)}
}

func (_c *UserStorage_Update_Call) Run(run func(user *domain.User)) *UserStorage_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.User))
	})
	return _c
}

func (_c *UserStorage_Update_Call) Return(_a0 error) *UserStorage_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewUserStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserStorage creates a new instance of UserStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserStorage(t mockConstructorTestingTNewUserStorage) *UserStorage {
	mock := &UserStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
