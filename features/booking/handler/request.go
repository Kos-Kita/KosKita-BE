package handler

import "KosKita/features/booking"

type BookRequest struct {
	BoardingHouseId uint   `json:"kos_id"`
	Method          string `json:"payment_type" form:"payment_type"`
	Bank            string `json:"bank" form:"bank"`
}

func RequestToCoreBook(input BookRequest, userIdLogin uint) booking.BookingCore {
	return booking.BookingCore{
		UserId: userIdLogin,
		BoardingHouseId: input.BoardingHouseId,
		Payment: booking.PaymentCore{
			Method: input.Method,
			Bank:   input.Bank,
		},
	}
}
