package handler

import "KosKita/features/chat"

type RoomRes struct {
	ID string `json:"room_id"`
}

type ChatRes struct {
	ID      string `json:"room_id"`
	Message string `json:"message"`
	UserID  uint   `json:"user_id"`
}

func CoreToGetChat(chat chat.CoreRoom) ChatRes {
	return ChatRes{
		ID:      chat.RoomID,
		Message: chat.Message,
		UserID:  chat.UserID,
	}
}

func CoreToGetChats(chats []chat.CoreRoom) []ChatRes {
	res := make([]ChatRes, 0)
	for _, chat := range chats {
		res = append(res, ChatRes{
			ID:      chat.RoomID,
			Message: chat.Message,
			UserID:  chat.UserID,
		})
	}
	return res
}

