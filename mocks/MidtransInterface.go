// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	booking "KosKita/features/booking"

	mock "github.com/stretchr/testify/mock"
)

// MidtransInterface is an autogenerated mock type for the MidtransInterface type
type MidtransInterface struct {
	mock.Mock
}

// CancelOrderPayment provides a mock function with given fields: bookingId
func (_m *MidtransInterface) CancelOrderPayment(bookingId string) error {
	ret := _m.Called(bookingId)

	if len(ret) == 0 {
		panic("no return value specified for CancelOrderPayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(bookingId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderPayment provides a mock function with given fields: book
func (_m *MidtransInterface) NewOrderPayment(book booking.BookingCore) (*booking.BookingCore, error) {
	ret := _m.Called(book)

	if len(ret) == 0 {
		panic("no return value specified for NewOrderPayment")
	}

	var r0 *booking.BookingCore
	var r1 error
	if rf, ok := ret.Get(0).(func(booking.BookingCore) (*booking.BookingCore, error)); ok {
		return rf(book)
	}
	if rf, ok := ret.Get(0).(func(booking.BookingCore) *booking.BookingCore); ok {
		r0 = rf(book)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*booking.BookingCore)
		}
	}

	if rf, ok := ret.Get(1).(func(booking.BookingCore) error); ok {
		r1 = rf(book)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMidtransInterface creates a new instance of MidtransInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMidtransInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MidtransInterface {
	mock := &MidtransInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
