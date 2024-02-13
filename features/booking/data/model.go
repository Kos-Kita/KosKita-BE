package data

import (
	"KosKita/features/booking"
	kd "KosKita/features/kos/data"
	ud "KosKita/features/user/data"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	Code            string  `gorm:"column:code; type:varchar(36);primary_key" json:"id"`
	Total           float64 `gorm:"column:total;"`
	UserId          uint
	BoardingHouseId uint
	BookedAt        time.Time        `gorm:"autoCreateTime"`
	Status          string           `gorm:"column:status; type:varchar(50);"`
	DeletedAt       gorm.DeletedAt   `gorm:"index"`
	User            ud.User          `gorm:"foreignKey:UserId"`
	BoardingHouse   kd.BoardingHouse `gorm:"foreignKey:BoardingHouseId"`
	Payment         Payment          `gorm:"embedded;embeddedPrefix:payment_"`
}

type Payment struct {
	Method        string     `gorm:"column:method; type:varchar(20);"`
	Bank          string     `gorm:"column:bank; type:varchar(20);"`
	VirtualNumber string     `gorm:"column:virtual_number; type:varchar(50);"`
	BillKey       string     `gorm:"column:bill_key; type:varchar(50);"`
	BillCode      string     `gorm:"column:bill_code; type:varchar(50);"`
	CreatedAt     time.Time  `gorm:"index"`
	ExpiredAt     *time.Time `gorm:"nullable"`
	PaidAt        *time.Time `gorm:"default:null;"`
}

func CoreToModelBook(input booking.BookingCore) Booking {
	return Booking{
		Code:            input.Code,
		UserId:          input.UserId,
		BoardingHouseId: input.BoardingHouseId,
		Total:           input.Total,
		BookedAt:        input.BookedAt,
		Status:          input.Status,
		Payment:         CoreToModelPay(input.Payment),
	}
}

func CoreToModelPay(input booking.PaymentCore) Payment {
	return Payment{
		Method:        input.Method,
		Bank:          input.Bank,
		VirtualNumber: input.VirtualNumber,
		BillKey:       input.BillKey,
		BillCode:      input.BillCode,
		ExpiredAt:     &input.ExpiredAt,
		PaidAt:        &input.PaidAt,
	}
}

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
		Payment:       PaymentModelToCore(model.Payment),
	}
}

func PaymentModelToCore(model Payment) booking.PaymentCore {
	return booking.PaymentCore{
		Method:        model.Method,
		Bank:          model.Bank,
		VirtualNumber: model.VirtualNumber,
		BillKey:       model.BillKey,
		BillCode:      model.BillCode,
		CreatedAt:     model.CreatedAt,
		ExpiredAt:     *model.ExpiredAt,
		PaidAt:        *model.PaidAt,
	}
}

func WebhoocksCoreToModel(reqNotif booking.BookingCore) Booking {
	return Booking{
		Code:   reqNotif.Code,
		Status: reqNotif.Status,
	}
}

// func (mod *Booking) GenerateCode() (err error) {
// 	// mod.Code, err = strconv.Atoi(fmt.Sprintf("%d%d%d", mod.UserId, mod.BoardingHouseId, time.Now().Unix()))
// 	var bookCode int
// 	mod.Code, err = strconv.Atoi(fmt.Sprintf("%d%d%d", mod.UserId, mod.BoardingHouseId, time.Now().Unix()))
// 	if err != nil {
// 		return err
// 	}
// 	// var stringCode string
// 	stringCode := strconv.Itoa(bookCode)
// 	mod.Code = stringCode

// 	return
// }
