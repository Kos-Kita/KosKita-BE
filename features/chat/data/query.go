package data

import (
	"KosKita/features/chat"

	"gorm.io/gorm"
)

type chatQuery struct {
	db   *gorm.DB
}

func New(db *gorm.DB) chat.ChatDataInterface {


	return &chatQuery{
		db: db,
	}
}

// CreateRoom implements chat.ChatDataInterface.
func (repo *chatQuery) CreateMessage(userIdLogin int, input chat.Core) (chat.Core, error) {
	chatInput := CoreToModelChat(input)
	chatInput.UserID = uint(userIdLogin)

	tx := repo.db.Create(&chatInput)
	if tx.Error != nil {
		return chat.Core{}, tx.Error
	}

	return chat.Core{
		ID:        chatInput.ID,
		Message:   chatInput.Message,
		RoomID:    chatInput.RoomID,
		UserID:    chatInput.UserID,
		CreatedAt: chatInput.CreatedAt,
		UpdatedAt: chatInput.UpdatedAt,
		User:      chatInput.User.ModelToCore(),
	}, nil
}
