// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	user "KosKita/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserServiceInterface is an autogenerated mock type for the UserServiceInterface type
type UserServiceInterface struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: userId, oldPassword, newPassword
func (_m *UserServiceInterface) ChangePassword(userId int, oldPassword string, newPassword string) error {
	ret := _m.Called(userId, oldPassword, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for ChangePassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, string) error); ok {
		r0 = rf(userId, oldPassword, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: input
func (_m *UserServiceInterface) Create(input user.Core) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: userId
func (_m *UserServiceInterface) Delete(userId int) error {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetById provides a mock function with given fields: userId
func (_m *UserServiceInterface) GetById(userId int) (*user.Core, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*user.Core, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int) *user.Core); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *UserServiceInterface) Login(email string, password string) (*user.Core, string, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *user.Core
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*user.Core, string, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *user.Core); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: userId, input
func (_m *UserServiceInterface) Update(userId int, input user.CoreUpdate) error {
	ret := _m.Called(userId, input)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, user.CoreUpdate) error); ok {
		r0 = rf(userId, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserServiceInterface creates a new instance of UserServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceInterface {
	mock := &UserServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
