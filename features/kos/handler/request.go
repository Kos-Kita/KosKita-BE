package handler

import (
	"KosKita/features/kos"
	"mime/multipart"
)

type KosRequest struct {
	Name          string   `json:"kos_name"`
	Description   string   `json:"description" form:"description"`
	Category      string   `json:"category"`
	Price         int      `json:"price"`
	Rooms         int      `json:"rooms"`
	Address       string   `json:"address"`
	Longitude     string   `json:"longitude"`
	Latitude      string   `json:"latitude"`
	KosFacilities []string `json:"kos_facilities" `
	KosRules      []string `json:"kos_rules"`
	UserID        uint
}

type KosFotoRequest struct {
	PhotoMain       *multipart.FileHeader `form:"main_kos_photo"`
	PhotoFront      *multipart.FileHeader ` form:"front_kos_photo"`
	PhotoBack       *multipart.FileHeader `form:"back_kos_photo"`
	PhotoRoomFront  *multipart.FileHeader ` form:"front_room_photo"`
	PhotoRoomInside *multipart.FileHeader ` form:"inside_room_photo"`
	UserID          uint
}

type RatingRequest struct {
	Score int `json:"score" form:"score"`
}

func RequestToCore(input KosRequest, userIdLogin uint) kos.CoreInput {
	var kosFacilities []kos.KosFacilityCore
	for _, facility := range input.KosFacilities {
		kosFacilities = append(kosFacilities, kos.KosFacilityCore{
			Facility: facility,
		})
	}

	var kosRules []kos.KosRuleCore
	for _, rule := range input.KosRules {
		kosRules = append(kosRules, kos.KosRuleCore{
			Rule: rule,
		})
	}

	return kos.CoreInput{
		UserID:        userIdLogin,
		Name:          input.Name,
		Description:   input.Description,
		Category:      input.Category,
		Price:         input.Price,
		Rooms:         input.Rooms,
		Address:       input.Address,
		Longitude:     input.Longitude,
		Latitude:      input.Latitude,
		KosFacilities: kosFacilities,
		KosRules:      kosRules,
	}
}

func RequestToCoreFoto(imageURLs []string, userIdLogin uint) kos.CoreFoto {
	return kos.CoreFoto{
		UserID:          userIdLogin,
		PhotoMain:       imageURLs[0],
		PhotoFront:      imageURLs[1],
		PhotoBack:       imageURLs[2],
		PhotoRoomFront:  imageURLs[3],
		PhotoRoomInside: imageURLs[4],
	}
}

func RequestToCorePut(input KosRequest, imageURLs []string, userIdLogin uint) kos.Core {
	kos := kos.Core{
		UserID:      userIdLogin,
		Name:        input.Name,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
		Rooms:       input.Rooms,
		Address:     input.Address,
		Longitude:   input.Longitude,
		Latitude:    input.Latitude,
		// KosFacilities: input.KosFacilities,
		// KosRules:      input.KosRules,
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
