package handler

import (
	"KosKita/features/booking"
)

type BookRequest struct {
	BoardingHouseId uint   `json:"kos_id"`
	Method          string `json:"payment_type" form:"payment_type"`
	Bank            string `json:"bank" form:"bank"`
}

type CancelBookingRequest struct {
	Status string `json:"status"`
}

func RequestToCoreBook(input BookRequest, userIdLogin uint) booking.BookingCore {
	return booking.BookingCore{
		UserId:          userIdLogin,
		BoardingHouseId: input.BoardingHouseId,
		Payment: booking.PaymentCore{
			Method: input.Method,
			Bank:   input.Bank,
		},
	}
}

func CancelRequestToCoreBooking(input CancelBookingRequest) booking.BookingCore {
	return booking.BookingCore{
		Payment: booking.PaymentCore{
			Status: input.Status,
		},
	}
}

type WebhoocksRequest struct {
	OrderID           int `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
}

func WebhoocksRequestToCore(input WebhoocksRequest) booking.BookingCore {
	return booking.BookingCore{
		Code:   input.OrderID,
		Status: input.TransactionStatus,
	}
}
