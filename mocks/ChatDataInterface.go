// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	chat "KosKita/features/chat"

	mock "github.com/stretchr/testify/mock"
)

// ChatDataInterface is an autogenerated mock type for the ChatDataInterface type
type ChatDataInterface struct {
	mock.Mock
}

// CreateMessage provides a mock function with given fields: userIdLogin, input
func (_m *ChatDataInterface) CreateMessage(userIdLogin int, input chat.Core) (chat.Core, error) {
	ret := _m.Called(userIdLogin, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessage")
	}

	var r0 chat.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(int, chat.Core) (chat.Core, error)); ok {
		return rf(userIdLogin, input)
	}
	if rf, ok := ret.Get(0).(func(int, chat.Core) chat.Core); ok {
		r0 = rf(userIdLogin, input)
	} else {
		r0 = ret.Get(0).(chat.Core)
	}

	if rf, ok := ret.Get(1).(func(int, chat.Core) error); ok {
		r1 = rf(userIdLogin, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: roomId
func (_m *ChatDataInterface) GetMessage(roomId string) ([]chat.Core, error) {
	ret := _m.Called(roomId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessage")
	}

	var r0 []chat.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]chat.Core, error)); ok {
		return rf(roomId)
	}
	if rf, ok := ret.Get(0).(func(string) []chat.Core); ok {
		r0 = rf(roomId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]chat.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roomId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewChatDataInterface creates a new instance of ChatDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewChatDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ChatDataInterface {
	mock := &ChatDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}