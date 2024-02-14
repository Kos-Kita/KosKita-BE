// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	booking "KosKita/features/booking"

	mock "github.com/stretchr/testify/mock"
)

// BookDataInterface is an autogenerated mock type for the BookDataInterface type
type BookDataInterface struct {
	mock.Mock
}

// CancelBooking provides a mock function with given fields: userIdLogin, bookingId, bookingCore
func (_m *BookDataInterface) CancelBooking(userIdLogin int, bookingId string, bookingCore booking.BookingCore) error {
	ret := _m.Called(userIdLogin, bookingId, bookingCore)

	if len(ret) == 0 {
		panic("no return value specified for CancelBooking")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, booking.BookingCore) error); ok {
		r0 = rf(userIdLogin, bookingId, bookingCore)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBooking provides a mock function with given fields: userId
func (_m *BookDataInterface) GetBooking(userId uint) ([]booking.BookingCore, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetBooking")
	}

	var r0 []booking.BookingCore
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]booking.BookingCore, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint) []booking.BookingCore); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]booking.BookingCore)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalBooking provides a mock function with given fields:
func (_m *BookDataInterface) GetTotalBooking() (int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTotalBooking")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func() (int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalBookingPerMonth provides a mock function with given fields: year, month
func (_m *BookDataInterface) GetTotalBookingPerMonth(year int, month int) (int, error) {
	ret := _m.Called(year, month)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalBookingPerMonth")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (int, error)); ok {
		return rf(year, month)
	}
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(year, month)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(year, month)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: userIdLogin, input
func (_m *BookDataInterface) Insert(userIdLogin int, input booking.BookingCore) (*booking.BookingCore, error) {
	ret := _m.Called(userIdLogin, input)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 *booking.BookingCore
	var r1 error
	if rf, ok := ret.Get(0).(func(int, booking.BookingCore) (*booking.BookingCore, error)); ok {
		return rf(userIdLogin, input)
	}
	if rf, ok := ret.Get(0).(func(int, booking.BookingCore) *booking.BookingCore); ok {
		r0 = rf(userIdLogin, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*booking.BookingCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, booking.BookingCore) error); ok {
		r1 = rf(userIdLogin, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WebhoocksData provides a mock function with given fields: webhoocksReq
func (_m *BookDataInterface) WebhoocksData(webhoocksReq booking.BookingCore) error {
	ret := _m.Called(webhoocksReq)

	if len(ret) == 0 {
		panic("no return value specified for WebhoocksData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(booking.BookingCore) error); ok {
		r0 = rf(webhoocksReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBookDataInterface creates a new instance of BookDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookDataInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookDataInterface {
	mock := &BookDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
