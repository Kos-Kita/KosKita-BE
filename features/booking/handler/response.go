package handler

import (
	"KosKita/features/booking"
	"KosKita/features/kos"
	"time"
)

type BookingResponse struct {
	Code                 string     `json:"booking_code"`
	Status               string     `json:"status"`
	Total                float64    `json:"total"`
	PaymentMethod        string     `json:"payment_method"`
	PaymentBank          string     `json:"bank"`
	PaymentVirtualNumber string     `json:"virtual_number"`
	PaymentExpiredAt     *time.Time `json:"payment_expired"`
}

type BookingHistoryResponse struct {
	KosId         uint     `json:"kos_id"`
	KosName       string   `json:"kos_name"`
	KosFasilitas  []string `json:"kos_fasilitas"`
	KosLokasi     string   `json:"kos_lokasi"`
	KosRating     float64  `json:"kos_ratingy"`
	KosMainFoto   string   `json:"kos_main_foto"`
	PaymentStatus string   `json:"payment_status"`
	TotalHarga    float64  `json:"total_harga"`
}

func CoreToResponseBook(core *booking.BookingCore) BookingResponse {
	return BookingResponse{
		Code:                 core.Code,
		Status:               core.Status,
		Total:                core.Total,
		PaymentMethod:        core.Method,
		PaymentBank:          core.Bank,
		PaymentVirtualNumber: core.VirtualNumber,
		PaymentExpiredAt:     &core.ExpiredAt,
	}
}

func CoreToResponseBookHistory(core *booking.BookingCore) BookingHistoryResponse {
	return BookingHistoryResponse{
		KosId:         core.BoardingHouse.ID,
		KosName:       core.BoardingHouse.Name,
		KosLokasi:     core.BoardingHouse.Address,
		KosRating:     KosRatingResult(core.BoardingHouse.Ratings),
		KosMainFoto:   core.BoardingHouse.PhotoMain,
		PaymentStatus: core.Status,
		TotalHarga:    core.Total,
	}
}

func KosFasilitasList(kf []kos.KosFacilityCore) []string {
	var results []string
	for _, v := range kf {
		results = append(results, v.Facility)
	}
	return results
}

func KosRatingResult(numbers []kos.RatingCore) float64 {
	var results float64
	if len(numbers) > 0 {
		for _, num := range numbers {
			results += float64(num.Score)
		}
		return float64(results) / float64(len(numbers))
	}
	return 0
}
