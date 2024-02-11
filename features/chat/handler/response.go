package handler

type RoomRes struct {
	ID string `json:"room_id"`
}

type ChatRes struct {
	ID      string `json:"room_id"`
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}
