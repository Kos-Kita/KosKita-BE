package handler

import (
	"KosKita/features/kos"
	"mime/multipart"
)

type KosRequest struct {
	Name            string                `json:"kos_name" form:"kos_name"`
	Description     string                `json:"description" form:"description"`
	Category        string                `json:"category" form:"category"`
	Price           int                   `json:"price" form:"price"`
	Rooms           int                   `json:"rooms" form:"rooms"`
	Address         string                `json:"address" form:"address"`
	Longitude       string                `json:"longitude" form:"longitude"`
	Latitude        string                `json:"latitude" form:"latitude"`
	KosFacilities   string                `json:"kos_facilities" form:"kos_facilities"`
	KosRules        string                `json:"kos_rules" form:"kos_rules"`
	PhotoMain       *multipart.FileHeader `json:"main_kos_photo" form:"main_kos_photo"`
	PhotoFront      *multipart.FileHeader `json:"front_kos_photo" form:"front_kos_photo"`
	PhotoBack       *multipart.FileHeader `json:"back_kos_photo" form:"back_kos_photo"`
	PhotoRoomFront  *multipart.FileHeader `json:"front_room_photo" form:"front_room_photo"`
	PhotoRoomInside *multipart.FileHeader `json:"inside_room_photo" form:"inside_room_photo"`
	UserID          uint
}

type RatingRequest struct {
	Score int `json:"score" form:"score"`
}

func RequestToCore(input KosRequest, imageURLs []string, userIdLogin uint) kos.CoreInput {
	return kos.CoreInput{
		UserID:          userIdLogin,
		Name:            input.Name,
		Description:     input.Description,
		Category:        input.Category,
		Price:           input.Price,
		Rooms:           input.Rooms,
		Address:         input.Address,
		Longitude:       input.Longitude,
		Latitude:        input.Latitude,
		KosFacilities:   input.KosFacilities,
		KosRules:        input.KosRules,
		PhotoMain:       imageURLs[0],
		PhotoFront:      imageURLs[1],
		PhotoBack:       imageURLs[2],
		PhotoRoomFront:  imageURLs[3],
		PhotoRoomInside: imageURLs[4],
	}
}

func RequestToCorePut(input KosRequest, imageURLs []string, userIdLogin uint) kos.Core {
	kos := kos.Core{
		UserID:        userIdLogin,
		Name:          input.Name,
		Description:   input.Description,
		Category:      input.Category,
		Price:         input.Price,
		Rooms:         input.Rooms,
		Address:       input.Address,
		Longitude:     input.Longitude,
		Latitude:      input.Latitude,
		KosFacilities: input.KosFacilities,
		KosRules:      input.KosRules,
	}

	if len(imageURLs) >= 5 {
		kos.PhotoMain = imageURLs[0]
		kos.PhotoFront = imageURLs[1]
		kos.PhotoBack = imageURLs[2]
		kos.PhotoRoomFront = imageURLs[3]
		kos.PhotoRoomInside = imageURLs[4]
	}

	return kos
}

func RequestToCoreRating(input RatingRequest, kosId uint, userIdLogin uint) kos.RatingCore {
	return kos.RatingCore{
		UserID:          userIdLogin,
		BoardingHouseID: kosId,
		Score:           input.Score,
	}
}
