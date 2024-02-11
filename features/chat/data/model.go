package data

import (
	"KosKita/features/chat"
	"KosKita/features/user/data"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Message string
	RoomID  string
	UserID  uint
	User    data.User
}

func CoreToModelChat(input chat.Core) Chat {
	return Chat{
		Message: input.Message,
		RoomID:  input.RoomID,
		UserID:  input.UserID,
	}
}

func (m Chat) ModelToCoreChat() chat.Core {
	return chat.Core{
		Message:   m.Message,
		RoomID:    m.RoomID,
		UserID:    m.UserID,
	}
}
