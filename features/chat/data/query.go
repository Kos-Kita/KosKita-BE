package data

import (
	"KosKita/features/chat"

	"gorm.io/gorm"
)

type chatQuery struct {
	db *gorm.DB
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
		// ID:        chatInput.ID,
		Message: chatInput.Message,
		RoomID:  chatInput.RoomID,
		UserID:  chatInput.UserID,
		// CreatedAt: chatInput.CreatedAt,
		// UpdatedAt: chatInput.UpdatedAt,
		// User:      chatInput.User.ModelToCore(),
	}, nil
}

// GetMessage implements chat.ChatDataInterface.
func (repo *chatQuery) GetMessage(roomId string) ([]chat.Core, error) {
	var chats []Chat
	tx := repo.db.Where("room_id = ?", roomId).Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var cores []chat.Core
	for _, c := range chats {
		cores = append(cores, c.ModelToCoreChat())
	}

	return cores, nil
}
