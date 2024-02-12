package data

import (
	"KosKita/features/booking"
	kd "KosKita/features/kos/data"
	ud "KosKita/features/user/data"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	Code            string  `gorm:"column:code; primaryKey;"`
	Total           float64 `gorm:"column:total;"`
	UserId          uint
	BoardingHouseId uint
	BookedAt        time.Time        `gorm:"autoCreateTime"`
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
	Status        string     `gorm:"column:status; type:varchar(50);"`
	CreatedAt     time.Time  `gorm:"index"`
	ExpiredAt     *time.Time `gorm:"nullable"`
	PaidAt        *time.Time `gorm:"default:null;"`
}

func CoreToModelBook(input booking.BookingCore) Booking {
	return Booking{
		Code:            input.Code,
		UserId:          input.UserId,
		BoardingHouseId: input.BoardingHouseId,
	}
}

func CoreToModelBookCancel(input booking.BookingCore) Booking {
	return Booking{
		Payment: Payment{
			Status: input.Payment.Status,
		},
	}
}

func ModelToCoreBook(model Booking) booking.BookingCore {
	return booking.BookingCore{
		Code:          model.Code,
		Total:         model.Total,
		UserId:        model.UserId,
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
		Status:        model.Status,
		CreatedAt:     model.CreatedAt,
		ExpiredAt:     *model.ExpiredAt,
		PaidAt:        *model.PaidAt,
	}
}

func (mod *Booking) GenerateCode() (err error) {
	// mod.Code, err = strconv.Atoi(fmt.Sprintf("%d%d%d", mod.UserId, mod.BoardingHouseId, time.Now().Unix()))
	var bookCode int
	bookCode, err = strconv.Atoi(fmt.Sprintf("%d%d%d", mod.UserId, mod.BoardingHouseId, time.Now().Unix()))
	if err != nil {
		return err
	}
	// var stringCode string
	stringCode := strconv.Itoa(bookCode)
	mod.Code = stringCode

	return
}

func WebhoocksCoreToModel(reqNotif booking.BookingCore) Booking {
	return Booking{
		Code: reqNotif.Code,
		Payment: Payment{
			Status: reqNotif.Payment.Status,
		},
	}
}

// func WebhoocksCoreToModel(reqNotif booking.BookingCore) (Booking, error) {

// return Booking{
// 	Code:    reqNotif.Code,
// 	Payment: payment,
// }, nil
//	}

// func Notif(reqNotif booking.PaymentCore) (Payment, error) {
// 	switch reqNotif.Status {
// 	case "pending", "success", "failed", "expired":
// 		return Payment{
// 			Status: reqNotif.Status,
// 		}, nil
// 	default:
// 		return Payment{}, fmt.Errorf("invalid payment status: %s", reqNotif.Status)
// 	}
// }
