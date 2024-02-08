package handler

import (
	"KosKita/features/booking"
	"time"
)

type BookingResponse struct {
	Code                 int        `json:"booking_code,omitempty"`
	Status               string     `json:"status,omitempty"`
	Total                float64    `json:"total,omitempty"`
	PaymentBank          string     `json:"payment_method,omitempty"`
	PaymentVirtualNumber string     `json:"virtual_number,omitempty"`
	PaymentBillKey       string     `json:"key_bill,omitempty"`
	PaymentBillCode      string     `json:"code_bill,omitempty"`
	PaymentExpiredAt     *time.Time `json:"payment_expired,omitempty"`
}

func CoreToResponseBook(core *booking.BookingCore) BookingResponse {
	return BookingResponse{
		Code:        core.Code,
		Status:      core.Payment.Status,
		Total:       core.Total,
		PaymentBank: core.Payment.Bank,
		PaymentVirtualNumber: core.Payment.VirtualNumber,
		PaymentBillKey:       core.Payment.BillKey,
		PaymentBillCode:      core.Payment.BillCode,
		PaymentExpiredAt:     &core.Payment.ExpiredAt,
	}
}
