package service

import (
	"KosKita/features/booking"
)

type bookService struct {
	bookData booking.BookDataInterface
}

func New(repo booking.BookDataInterface) booking.BookServiceInterface {
	return &bookService{
		bookData: repo,
	}
}

// Create implements booking.BookServiceInterface.
func (bs *bookService) Create(userIdLogin int, input booking.BookingCore) (*booking.BookingCore, error) {
	// Insert booking into database
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
