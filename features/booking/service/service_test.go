package service

import (
	"KosKita/features/booking"
	"KosKita/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostBooking(t *testing.T) {
	repo := new(mocks.BookDataInterface)
	srv := NewBooking(repo)

	input := booking.BookingCore{
		ID:            "10843",
		UserID:          1,
		BoardingHouseId: 1,
		StartDate:       "2024-01-01",
		EndDate:         "2024-12-31",
		PaymentType:     "credit_card",
		Total:           100000,
		Bank:            "BCA",
	}

	// response := booking.BookingCore{
	// 	ID:            "10843",
	// 	StartDate:     "2024-01-01",
	// 	EndDate:       "2024-12-31",
	// 	PaymentType:   "credit_card",
	// 	Total:         100000,
	// 	Status:        "pending",
	// 	Bank:          "BCA",
	// 	VirtualNumber: "1234567890",
	// 	ExpiredAt:     "2024-12-31",
	// }

	t.Run("invalid user id", func(t *testing.T) {
		_, err := srv.PostBooking(0, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "invalid id")
	})

	t.Run("error from repository", func(t *testing.T) {
		repo.On("PostBooking", uint(1), input).Return(nil, errors.New("database error"))

		_, err := srv.PostBooking(1, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

}
