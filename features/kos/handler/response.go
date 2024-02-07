package handler

import "KosKita/features/kos"

type KosResponseRating struct {
	ID            uint              `json:"id" form:"id"`
	Name          string            `json:"kos_name" form:"kos_name"`
	Rating        int               `json:"rating" form:"rating"`
	Category      string            `json:"category" form:"category"`
	Price         int               `json:"price" form:"price"`
	Address       string            `json:"address" form:"address"`
	KosFacilities string            `json:"kos_facilities" form:"kos_facilities"`
	PhotoKos      PhotoMainResponse `json:"photo_kos" form:"photo_kos"`
}

type PhotoMainResponse struct {
	PhotoMain string `json:"main_kos_photo"`
}

func CoreToGetRating(rating kos.RatingCore) KosResponseRating {
	return KosResponseRating{
		ID:            rating.BoardingHouse.ID,
		Name:          rating.BoardingHouse.Name,
		Rating:        rating.Score,
		Category:      rating.BoardingHouse.Category,
		Price:         rating.BoardingHouse.Price,
		Address:       rating.BoardingHouse.Address,
		KosFacilities: rating.BoardingHouse.KosFacilities,
		PhotoKos:      PhotoMainResponse{PhotoMain: rating.BoardingHouse.PhotoMain},
	}
}

// type KosResponseRating struct {
// 	ID              uint   `json:"id" form:"id"`
// 	Name            string `json:"kos_name" form:"kos_name"`
// 	Description     string `json:"description" form:"description"`
// 	Category        string `json:"category" form:"category"`
// 	Price           int    `json:"price" form:"price"`
// 	Rooms           int    `json:"stock" form:"stock"`
// 	Address         string `json:"address" form:"address"`
// 	KosFacilities   string `json:"kos_facilities" form:"kos_facilities"`
// 	KosRules        string `json:"kos_rules" form:"kos_rules"`
// 	PhotoMain       string `json:"main_kos_photo" form:"main_kos_photo"`
// 	PhotoFront      string `json:"front_kos_photo" form:"front_kos_photo"`
// 	PhotoBack       string `json:"back_kos_photo" form:"back_kos_photo"`
// 	PhotoRoomFront  string `json:"front_room_photo" form:"front_room_photo"`
// 	PhotoRoomInside string `json:"inside_room_photo" form:"inside_room_photo"`
// }
