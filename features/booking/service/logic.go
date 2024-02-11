package service

import (
	"KosKita/features/booking"
	"KosKita/features/user"
	"errors"
)

type bookService struct {
	bookData    booking.BookDataInterface
	userService user.UserServiceInterface
}

func New(repo booking.BookDataInterface, us user.UserServiceInterface) booking.BookServiceInterface {
	return &bookService{
		bookData:    repo,
		userService: us,
	}
}

// Create implements booking.BookServiceInterface.
func (bs *bookService) Create(userIdLogin int, input booking.BookingCore) (*booking.BookingCore, error) {
	user, err := bs.userService.GetById(userIdLogin)
	if err != nil {
		return nil, err
	}

	if user.Role != "renter" {
		return nil, errors.New("anda bukan renter")
	}

	bookCore, err := bs.bookData.Insert(userIdLogin, input)
	if err != nil {
		return nil, err
	}

	return bookCore, nil
}

// CancelBooking implements booking.BookServiceInterface.
func (bs *bookService) CancelBooking(userIdLogin int, bookingId string, bookingCore booking.BookingCore) error {
	if bookingCore.Payment.Status == "" {
		bookingCore.Payment.Status = "cancelled"
	}

	err := bs.bookData.CancelBooking(userIdLogin, bookingId, bookingCore)
	return err
}
