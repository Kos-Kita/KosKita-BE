package chat

import (
	"KosKita/features/user"
	"time"
)

type Core struct {
	ID         uint
	Message    string
	RoomID     string
	ReceiverID uint
	SenderID   uint
	User       user.Core
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// interface untuk Data Layer
type ChatDataInterface interface {
	CreateRoom(roomID string, receiverID int, senderID int) error
	CreateMessage(userIdLogin int, input Core) (Core, error)
	GetMessage(roomId string) ([]Core, error)
}

// interface untuk Service Layer
type ChatServiceInterface interface {
	CreateRoom(roomID string, receiverID int, senderID int) error
	CreateChat(userIdLogin int, input Core) (Core, error)
	GetMessage(roomId string) ([]Core, error)
}
