package handler

import (
	"KosKita/features/kos"
	"KosKita/features/user/handler"
)

type KosResponseRating struct {
	ID            uint              `json:"id" form:"id"`
	Name          string            `json:"kos_name" form:"kos_name"`
	Rating        float64            `json:"rating" form:"rating"`
	Category      string            `json:"category" form:"category"`
	Price         int               `json:"price" form:"price"`
	Address       string            `json:"address" form:"address"`
	KosFacilities string            `json:"kos_facilities" form:"kos_facilities"`
	PhotoKos      PhotoMainResponse `json:"photo_kos" form:"photo_kos"`
}

type KosResponseUser struct {
	ID            uint              `json:"id" form:"id"`
	Name          string            `json:"kos_name" form:"kos_name"`
	Rating        float64             `json:"rating" form:"rating"`
	Address       string            `json:"address" form:"address"`
	KosFacilities string            `json:"kos_facilities" form:"kos_facilities"`
	PhotoKos      PhotoMainResponse `json:"photo_kos" form:"photo_kos"`
}

type PhotoMainResponse struct {
	PhotoMain string `json:"main_kos_photo"`
}

type KosResponseDetail struct {
	ID            uint                          `json:"id" form:"id"`
	Name          string                        `json:"kos_name" form:"kos_name"`
	Description   string                        `json:"description" form:"description"`
	Category      string                        `json:"category" form:"category"`
	Rating        float64                          `json:"rating" form:"rating"`
	Price         int                           `json:"price" form:"price"`
	Rooms         int                           `json:"stock" form:"stock"`
	Address       string                        `json:"address" form:"address"`
	KosFacilities string                        `json:"kos_facilities" form:"kos_facilities"`
	KosRules      string                        `json:"kos_rules" form:"kos_rules"`
	PhotoKos      PhotoDetailResponse           `json:"photo_kos" form:"photo_kos"`
	User          handler.UserKosDetailResponse `json:"user" form:"user"`
}

type PhotoDetailResponse struct {
	PhotoMain       string `json:"main_kos_photo" form:"main_kos_photo"`
	PhotoFront      string `json:"front_kos_photo" form:"front_kos_photo"`
	PhotoBack       string `json:"back_kos_photo" form:"back_kos_photo"`
	PhotoRoomFront  string `json:"front_room_photo" form:"front_room_photo"`
	PhotoRoomInside string `json:"inside_room_photo" form:"inside_room_photo"`
}

func CoreToGetRating(kos kos.Core) KosResponseRating {
	var totalRating float64
	for _, rating := range kos.Ratings {
		totalRating += float64(rating.Score)
	}
	
	var averageRating float64
	if len(kos.Ratings) > 0 {
		averageRating = totalRating / float64(len(kos.Ratings))
	}

	return KosResponseRating{
		ID:            kos.ID,
		Name:          kos.Name,
		Rating:        averageRating,
		Category:      kos.Category,
		Price:         kos.Price,
		Address:       kos.Address,
		KosFacilities: kos.KosFacilities,
		PhotoKos:      PhotoMainResponse{PhotoMain: kos.PhotoMain},
	}
}

func CoreToGetDetail(kos kos.Core) KosResponseDetail {
	var totalRating float64
	for _, rating := range kos.Ratings {
		totalRating += float64(rating.Score)
	}
	
	var averageRating float64
	if len(kos.Ratings) > 0 {
		averageRating = totalRating / float64(len(kos.Ratings))
	}

	return KosResponseDetail{
		ID:            kos.ID,
		Name:          kos.Name,
		Description:   kos.Description,
		Rooms:         kos.Rooms,
		Rating:        averageRating,
		Category:      kos.Category,
		Price:         kos.Price,
		Address:       kos.Address,
		KosFacilities: kos.KosFacilities,
		KosRules:      kos.KosRules,
		PhotoKos: PhotoDetailResponse{
			PhotoMain:       kos.PhotoMain,
			PhotoFront:      kos.PhotoFront,
			PhotoBack:       kos.PhotoBack,
			PhotoRoomFront:  kos.PhotoRoomFront,
			PhotoRoomInside: kos.PhotoRoomInside,
		},
		User: handler.UserKosDetailResponse{
			Name:         kos.User.Name,
			UserName:     kos.User.UserName,
			PhotoProfile: kos.User.PhotoProfile,
		},
	}
}

func CoreToGetUser(kos kos.Core) KosResponseUser {
	var totalRating float64
	for _, rating := range kos.Ratings {
		totalRating += float64(rating.Score)
	}
	
	var averageRating float64
	if len(kos.Ratings) > 0 {
		averageRating = totalRating / float64(len(kos.Ratings))
	}

	return KosResponseUser{
		ID:            kos.ID,
		Name:          kos.Name,
		Rating:        averageRating,
		Address:       kos.Address,
		KosFacilities: kos.KosFacilities,
		PhotoKos:      PhotoMainResponse{PhotoMain: kos.PhotoMain},
	}
}