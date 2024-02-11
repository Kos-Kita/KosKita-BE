package data

import (
	"KosKita/features/chat"
	"KosKita/features/user/data"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Message string
	RoomID  string
	UserID  uint
	User    data.User
}

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Chat
	ID       string
	RoomID   string
	Username string
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
		ID:        m.ID,
		Message:   m.Message,
		RoomID:    m.RoomID,
		UserID:    m.UserID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		User:      m.User.ModelToCore(),
	}
}
