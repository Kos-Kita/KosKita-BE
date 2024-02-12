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
func (repo *chatQuery) CreateMessage(userIdLogin int, input chat.Core) (chat.CoreRoom, error) {
	chatInput := CoreToModelChat(input)
	chatInput.UserID = uint(userIdLogin)

	tx := repo.db.Create(&chatInput)
	if tx.Error != nil {
		return chat.CoreRoom{}, tx.Error
	}

	return chat.CoreRoom{
		Message: chatInput.Message,
		RoomID:  chatInput.RoomID,
		UserID:  chatInput.UserID,
	}, nil
}

// GetMessage implements chat.ChatDataInterface.
func (repo *chatQuery) GetMessage(roomId string) ([]chat.CoreRoom, error) {
	var chats []Chat
	tx := repo.db.Where("room_id = ?", roomId).Find(&chats)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var cores []chat.CoreRoom
	for _, c := range chats {
		cores = append(cores, c.ModelToCoreChat())
	}

	return cores, nil
}
