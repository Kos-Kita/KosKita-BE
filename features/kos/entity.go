package kos

import ("time")
// "KosKita/features/user")

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
	// User user.Core
}

type KosDataInterface interface {
	Insert(userIdLogin int, input Core) error
}

// interface untuk Service Layer
type KosServiceInterface interface {
	Create(userIdLogin int, input Core) error
}
