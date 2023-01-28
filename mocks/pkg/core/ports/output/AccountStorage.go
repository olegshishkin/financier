// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/olegshishkin/financier/pkg/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// AccountStorage is an autogenerated mock type for the AccountStorage type
type AccountStorage struct {
	mock.Mock
}

type AccountStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountStorage) EXPECT() *AccountStorage_Expecter {
	return &AccountStorage_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: account
func (_m *AccountStorage) Create(account *domain.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountStorage_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AccountStorage_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - account *domain.Account
func (_e *AccountStorage_Expecter) Create(account interface{}) *AccountStorage_Create_Call {
	return &AccountStorage_Create_Call{Call: _e.mock.On("Create", account)}
}

func (_c *AccountStorage_Create_Call) Run(run func(account *domain.Account)) *AccountStorage_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*domain.Account))
	})
	return _c
}

func (_c *AccountStorage_Create_Call) Return(_a0 error) *AccountStorage_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindEnabledByName provides a mock function with given fields: name
func (_m *AccountStorage) FindEnabledByName(name string) (*domain.Account, error) {
	ret := _m.Called(name)

	var r0 *domain.Account
	if rf, ok := ret.Get(0).(func(string) *domain.Account); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountStorage_FindEnabledByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindEnabledByName'
type AccountStorage_FindEnabledByName_Call struct {
	*mock.Call
}

// FindEnabledByName is a helper method to define mock.On call
//   - name string
func (_e *AccountStorage_Expecter) FindEnabledByName(name interface{}) *AccountStorage_FindEnabledByName_Call {
	return &AccountStorage_FindEnabledByName_Call{Call: _e.mock.On("FindEnabledByName", name)}
}

func (_c *AccountStorage_FindEnabledByName_Call) Run(run func(name string)) *AccountStorage_FindEnabledByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AccountStorage_FindEnabledByName_Call) Return(_a0 *domain.Account, _a1 error) *AccountStorage_FindEnabledByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewAccountStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountStorage creates a new instance of AccountStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountStorage(t mockConstructorTestingTNewAccountStorage) *AccountStorage {
	mock := &AccountStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}