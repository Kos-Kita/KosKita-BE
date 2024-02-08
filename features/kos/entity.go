package kos

import (
	"time"
	"KosKita/features/user"
)

type Core struct {
	ID              uint
	Name            string `validate:"required"`
	Description     string
	Category        string `validate:"oneof=putra putri campur"`
	Price           int    `validate:"required"`
	Rooms           int    `validate:"required"`
	Address         string `validate:"required"`
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
}

type RatingCore struct {
	ID              uint
	Score           int
	UserID          uint
	BoardingHouseID uint
	BoardingHouse   Core
	CreatedAt       time.Time
	UpdatedAt       time.Time
	User user.Core
}

type KosDataInterface interface {
	Insert(userIdLogin int, input Core) error
	Update(userIdLogin int, input Core) error
	InsertRating(userIdLogin, kosId int, score RatingCore) error
	SelectByRating() ([]RatingCore, error)
	Delete(userIdLogin, kosId int) error
	SelectById(kosId int)(*RatingCore, error)
	SelectByUserId(userIdLogin int)([]RatingCore, error)
}

// interface untuk Service Layer
type KosServiceInterface interface {
	Create(userIdLogin int, input Core) error
	Put(userIdLogin int, input Core) error
	CreateRating(userIdLogin, kosId int, score RatingCore) error
	GetByRating() ([]RatingCore, error)
	Delete(userIdLogin, kosId int) error
	GetById(kosId int)(*RatingCore, error)
	GetByUserId(userIdLogin int)([]RatingCore, error)
}
