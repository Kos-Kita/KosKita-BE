package chat

import (
	"KosKita/features/user"
	"time"
)

type Core struct {
	ID        uint
	Message   string
	RoomID    string
	UserID    uint
	User      user.Core
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CoreRoom struct {
	Message   string
	RoomID    string
	UserID    uint
}

// interface untuk Data Layer
type ChatDataInterface interface {
	CreateMessage(userIdLogin int, input Core) (CoreRoom, error)
	GetMessage(roomId string) ([]CoreRoom, error)
}

// interface untuk Service Layer
type ChatServiceInterface interface {
	CreateChat(userIdLogin int, input Core) (CoreRoom, error)
	GetMessage(roomId string) ([]CoreRoom, error)
}
