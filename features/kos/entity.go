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
	Ratings []RatingCore
	User user.Core
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
	Insert(userIdLogin int, input Core) error
	Update(userIdLogin int, input Core) error
	CekRating(userId, kosId int) (*RatingCore, error)
	InsertRating(userIdLogin, kosId int, score RatingCore) error
	SelectByRating() ([]Core, error)
	Delete(userIdLogin, kosId int) error
	SelectById(kosId int)(*RatingCore, error)
	SelectByUserId(userIdLogin int)([]RatingCore, error)
}

// interface untuk Service Layer
type KosServiceInterface interface {
	Create(userIdLogin int, input Core) error
	Put(userIdLogin int, input Core) error
	CreateRating(userIdLogin, kosId int, score RatingCore) error
	GetByRating() ([]Core, error)
	Delete(userIdLogin, kosId int) error
	GetById(kosId int)(*RatingCore, error)
	GetByUserId(userIdLogin int)([]RatingCore, error)
}
