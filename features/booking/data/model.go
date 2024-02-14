package data

import (
	"KosKita/features/booking"
	kd "KosKita/features/kos/data"
	ud "KosKita/features/user/data"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	Code            string `gorm:"type:varchar(36);primary_key" json:"id"`
	Total           float64
	UserId          uint
	BoardingHouseId uint
	Status          string
	BookedAt        time.Time        `gorm:"autoCreateTime"`
	DeletedAt       gorm.DeletedAt   `gorm:"index"`
	CreatedAt       time.Time        `gorm:"index"`
	User            ud.User          `gorm:"foreignKey:UserId"`
	BoardingHouse   kd.BoardingHouse `gorm:"foreignKey:BoardingHouseId"`
	Method          string           `gorm:"column:method; type:varchar(20);"`
	Bank            string
	VirtualNumber   string `gorm:"column:virtual_number; type:varchar(50);"`
	ExpiredAt       time.Time
	UpdatedAt       time.Time
}

func CoreToModelBook(input booking.BookingCore) Booking {
	return Booking{
		Code:            input.Code,
		UserId:          input.UserId,
		BoardingHouseId: input.BoardingHouseId,
		Total:           input.Total,
		BookedAt:        input.BookedAt,
		Status:          input.Status,
		Method:          input.Method,
		Bank:            input.Bank,
		VirtualNumber:   input.VirtualNumber,
		ExpiredAt:       input.ExpiredAt,
		// Payment:         CoreToModelPay(input.Payment),
	}
}

// func CoreToModelPay(input booking.PaymentCore) Payment {
// 	return Payment{
// 		Method:        input.Method,
// 		Bank:          input.Bank,
// 		VirtualNumber: input.VirtualNumber,
// 		BillKey:       input.BillKey,
// 		BillCode:      input.BillCode,
// 		ExpiredAt:     &input.ExpiredAt,
// 		PaidAt:        &input.PaidAt,
// 	}
// }

func CoreToModelBookCancel(input booking.BookingCore) Booking {
	return Booking{
		Status: input.Status,
	}
}

func ModelToCoreBook(model Booking) booking.BookingCore {
	return booking.BookingCore{
		Code:          model.Code,
		Total:         model.Total,
		UserId:        model.UserId,
		Status:        model.Status,
		BoardingHouse: model.BoardingHouse.ModelToCoreKos(),
		Method:        model.Method,
		Bank:          model.Bank,
		VirtualNumber: model.VirtualNumber,
		CreatedAt:     model.CreatedAt,
		ExpiredAt:     model.ExpiredAt,
		// Payment:       PaymentModelToCore(model.Payment),
	}
}

// func PaymentModelToCore(model Payment) booking.PaymentCore {
// 	return booking.PaymentCore{
// 		Method:        model.Method,
// 		Bank:          model.Bank,
// 		VirtualNumber: model.VirtualNumber,
// 		BillKey:       model.BillKey,
// 		BillCode:      model.BillCode,
// 		CreatedAt:     model.CreatedAt,
// 		ExpiredAt:     *model.ExpiredAt,
// 		PaidAt:        *model.PaidAt,
// 	}
// }

func WebhoocksCoreToModel(reqNotif booking.BookingCore) Booking {
	return Booking{
		Code:   reqNotif.Code,
		Status: reqNotif.Status,
	}
}
