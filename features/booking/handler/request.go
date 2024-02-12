package handler

import (
	"KosKita/features/booking"
	"strconv"
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

		Status: input.Status,
	}
}

type WebhoocksRequest struct {
	Code   string `json:"order_id"`
	Status string `json:"transaction_status"`
}

func WebhoocksRequestToCore(input WebhoocksRequest) booking.BookingCore {
	code, _ := strconv.Atoi(input.Code)
	return booking.BookingCore{
		Code: code,
		Status: input.Status,
	}
}
