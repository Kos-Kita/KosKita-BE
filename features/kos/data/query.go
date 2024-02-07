package data

import (
	"KosKita/features/kos"

	"gorm.io/gorm"
)

type kosQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) kos.KosDataInterface {
	return &kosQuery{
		db: db,
	}
}

// Insert implements kos.KosDataInterface.
func (repo *kosQuery) Insert(userIdLogin int, input kos.Core) error {
	kosInput := CoreToModel(input)
	kosInput.UserID = uint(userIdLogin)
	
	tx := repo.db.Create(&kosInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Update implements kos.KosDataInterface.
func (repo *kosQuery) Update(userIdLogin int, input kos.Core) error {
	kos := BoardingHouse{}
	tx := repo.db.Where("id = ? AND user_id = ?", input.ID, userIdLogin).First(&kos)
	if tx.Error != nil {
		return tx.Error
	}

	kosInput := CoreToModel(input)

	tx = repo.db.Model(&kos).Updates(&kosInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}