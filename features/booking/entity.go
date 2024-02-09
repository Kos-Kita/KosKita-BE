package booking

import (
	"time"

	kd "KosKita/features/kos"
	ud "KosKita/features/user"

	"gorm.io/gorm"
)

type BookingCore struct {
	Code            int
	Total           float64
	BookedAt        time.Time
	DeletedAt       gorm.DeletedAt
	UserId          uint
	User            ud.Core
	BoardingHouseId uint
	BoardingHouse   kd.Core
	Payment PaymentCore
}

type PaymentCore struct {
	Method        string
	Bank          string
	VirtualNumber string
	BillKey       string
	BillCode      string
	Status        string
	BookingCode   int
	BookingTotal  float64
	CreatedAt     time.Time
	ExpiredAt     time.Time
	PaidAt        time.Time
}

type BookDataInterface interface {
	Insert(userIdLogin int, input BookingCore) (*BookingCore, error)
}

// interface untuk Service Layer
type BookServiceInterface interface {
	Create(userIdLogin int, input BookingCore) (*BookingCore, error)
}