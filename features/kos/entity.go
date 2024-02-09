package kos

import (
	"KosKita/features/user"
	"time"
)

type Core struct {
	ID              uint
	Name            string `validate:"required"`
	Description     string
	Category        string `validate:"oneof=putra putri campur"`
	Price           int    `validate:"required"`
	Rooms           int    `validate:"required"`
	Address         string `validate:"required"`
	Longitude       string
	Latitude        string
	KosFacilities   string
	KosRules        string
	PhotoMain       string
	PhotoFront      string
	PhotoBack       string
	PhotoRoomFront  string
	PhotoRoomInside string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UserID          uint
	Ratings         []RatingCore
	User            user.Core
	// User User
}

type CoreInput struct {
	ID              uint
	Name            string `validate:"required"`
	Description     string
	Category        string `validate:"oneof=putra putri campur"`
	Price           int    `validate:"required"`
	Rooms           int    `validate:"required"`
	Address         string `validate:"required"`
	Longitude       string
	Latitude        string
	KosFacilities   string
	KosRules        string
	PhotoMain       string
	PhotoFront      string
	PhotoBack       string
	PhotoRoomFront  string
	PhotoRoomInside string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	UserID          uint
	Ratings         []RatingCore
}

type User struct {
	ID           uint
	Name         string
	UserName     string
	Email        string
	Password     string
	Gender       string
	Role         string
	PhotoProfile string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RatingCore struct {
	ID              uint
	Score           int
	UserID          uint
	BoardingHouseID uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type KosDataInterface interface {
	Insert(userIdLogin int, input CoreInput) error
	Update(userIdLogin int, input Core) error
	CekRating(userId, kosId int) (*RatingCore, error)
	InsertRating(userIdLogin, kosId int, score RatingCore) error
	SelectByRating() ([]Core, error)
	Delete(userIdLogin, kosId int) error
	SelectById(kosId int) (*Core, error)
	SelectByUserId(userIdLogin int) ([]Core, error)
	SearchKos(query, category string, minPrice, maxPrice int) ([]Core, error)
}

// interface untuk Service Layer
type KosServiceInterface interface {
	Create(userIdLogin int, input CoreInput) error
	Put(userIdLogin int, input Core) error
	CreateRating(userIdLogin, kosId int, score RatingCore) error
	GetByRating() ([]Core, error)
	Delete(userIdLogin, kosId int) error
	GetById(kosId int) (*Core, error)
	GetByUserId(userIdLogin int) ([]Core, error)
	SearchKos(query, category string, minPrice, maxPrice int) ([]Core, error)
}
