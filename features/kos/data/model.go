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
	User            data.User
	Ratings         []Rating
}

type Rating struct {
	gorm.Model
	Score           int
	UserID          uint
	User            data.User
	BoardingHouseID uint
	BoardingHouse   BoardingHouse
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

func CoreToModelRating(input kos.RatingCore) Rating {
	return Rating{
		Score:           input.Score,
		UserID:          input.UserID,
		BoardingHouseID: input.BoardingHouseID,
	}
}

func (bh BoardingHouse) ModelToCoreKos() kos.Core {
	var ratings []kos.RatingCore
	for _, r := range bh.Ratings {
		ratings = append(ratings, r.ModelToCoreRating())
	}

	return kos.Core{
		UserID:          bh.UserID,
		ID:              bh.ID,
		Name:            bh.Name,
		Description:     bh.Description,
		Category:        bh.Category,
		Price:           bh.Price,
		Rooms:           bh.Rooms,
		Address:         bh.Address,
		KosFacilities:   bh.KosFacilities,
		KosRules:        bh.KosRules,
		PhotoMain:       bh.PhotoMain,
		PhotoFront:      bh.PhotoFront,
		PhotoBack:       bh.PhotoBack,
		PhotoRoomFront:  bh.PhotoRoomFront,
		PhotoRoomInside: bh.PhotoRoomInside,
		CreatedAt:       bh.CreatedAt,
		UpdatedAt:       bh.UpdatedAt,
		Ratings:         ratings,
		User:            bh.User.ModelToCore(),
	}
}


func (r Rating) ModelToCoreRating() kos.RatingCore {
	return kos.RatingCore{
		ID:              r.ID,
		Score:           r.Score,
		UserID:          r.UserID,
		BoardingHouseID: r.BoardingHouseID,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}
}

// func ModelToCoreKosRating(bh BoardingHouse, r Rating) kos.RatingCore {
// 	return kos.RatingCore{
// 		BoardingHouse: bh.ModelToCoreKos(),
// 		Score: r.ModelToCoreRating().Score,
// 	}
// }
