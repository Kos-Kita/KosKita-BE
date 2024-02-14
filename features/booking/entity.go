package booking

import (
	"time"

	kd "KosKita/features/kos"
	ud "KosKita/features/user"
)

type BookingCore struct {
	Code            string
	Total           float64
	BookedAt        time.Time
	DeletedAt       time.Time
	UserId          uint
	User            ud.Core
	BoardingHouseId uint
	BoardingHouse   kd.Core
	Method          string
	Bank            string
	VirtualNumber   string
	BookingCode     int
	BookingTotal    float64
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	ExpiredAt       time.Time
}

type MonthCount struct {
	Month int
	Count int
}

type BookDataInterface interface {
	Insert(userIdLogin int, input BookingCore) (*BookingCore, error)
	CancelBooking(userIdLogin int, bookingId string, bookingCore BookingCore) error
	GetBooking(userId uint) ([]BookingCore, error)
	WebhoocksData(webhoocksReq BookingCore) error
	GetTotalBooking() (int, error)
	GetTotalBookingPerMonth(year int, month int) (int, error)
}

// interface untuk Service Layer
type BookServiceInterface interface {
	Create(userIdLogin int, input BookingCore) (*BookingCore, error)
	CancelBooking(userIdLogin int, bookingId string, bookingCore BookingCore) error
	GetBooking(userId uint) ([]BookingCore, error)
	WebhoocksData(webhoocksReq BookingCore) error
}
