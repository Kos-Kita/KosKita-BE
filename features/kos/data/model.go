package data

import (
	"KosKita/features/kos"
"KosKita/features/user/data"
	"gorm.io/gorm"
)

type BoardingHouse struct {
	gorm.Model
	Name            string
	Description     string
	Category        string `gorm:"column:status; type:enum('putra', 'putri', 'campur');"`
	Price           int
	Rooms           int
	Address         string
	KosFacilities   string
	KosRules        string
	PhotoMain       string
	PhotoFront      string
	PhotoBack       string
	PhotoRoomFront  string
	PhotoRoomInside string
	UserID          uint
	User data.User
}

func CoreToModel(input kos.Core) BoardingHouse {
	return BoardingHouse{
		UserID:          input.UserID,
		Name:            input.Name,
		Description:     input.Description,
		Category:        input.Category,
		Price:           input.Price,
		Rooms:           input.Rooms,
		Address:         input.Address,
		KosFacilities:   input.KosFacilities,
		KosRules:        input.KosRules,
		PhotoMain:       input.PhotoMain,
		PhotoFront:      input.PhotoFront,
		PhotoBack:       input.PhotoBack,
		PhotoRoomFront:  input.PhotoRoomFront,
		PhotoRoomInside: input.PhotoRoomInside,
	}
}

func (bh BoardingHouse) ModelToCore() kos.Core {
	return kos.Core{
		UserID:          bh.UserID,
		ID:              bh.ID,
		Name:            bh.Name,
		Description:     bh.Description,
		Category:        bh.Category,
		Price:           bh.Price,
		Rooms:           bh.Rooms,
		Address:         bh.Address,
		PhotoMain:       bh.PhotoMain,
		PhotoFront:      bh.PhotoFront,
		PhotoBack:       bh.PhotoBack,
		PhotoRoomFront:  bh.PhotoRoomFront,
		PhotoRoomInside: bh.PhotoRoomInside,
		CreatedAt:       bh.CreatedAt,
		UpdatedAt:       bh.UpdatedAt,
	}
}
