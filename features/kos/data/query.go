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

// InsertRating implements kos.KosDataInterface.
func (repo *kosQuery) InsertRating(userIdLogin int, kosId int, score kos.RatingCore) error {
	ratingInput := CoreToModelRating(score)
	ratingInput.UserID = uint(userIdLogin)
	ratingInput.BoardingHouseID = uint(kosId)

	tx := repo.db.Create(&ratingInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectByRating implements kos.KosDataInterface.
func (repo *kosQuery) SelectByRating() ([]kos.Core, error) {
	var kosData []BoardingHouse
	var result []kos.Core

	// Menambahkan Preload untuk mengisi relasi User dan Ratings
	tx := repo.db.Preload("User").Preload("Ratings").Table("boarding_houses").
		Joins("left join ratings on boarding_houses.id = ratings.boarding_house_id").
		Group("boarding_houses.id").
		Select("boarding_houses.*, AVG(ratings.score) as average_rating").
		Order("average_rating desc").
		Find(&kosData)

	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, k := range kosData {
		result = append(result, k.ModelToCoreKos())
	}

	return result, nil
}

// Delete implements kos.KosDataInterface.
func (repo *kosQuery) Delete(userIdLogin int, kosId int) error {
	tx := repo.db.Where("id = ? AND user_id = ?", kosId, userIdLogin).Delete(&BoardingHouse{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectById implements kos.KosDataInterface.
func (repo *kosQuery) SelectById(kosId int) (*kos.RatingCore, error) {
	var kosDataGorm Rating
	tx := repo.db.Preload("User").Preload("BoardingHouse").Where("id = ?", kosId).First(&kosDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := kosDataGorm.ModelToCoreRating()
	return &result, nil
}

// SelectByUserId implements kos.KosDataInterface.
func (repo *kosQuery) SelectByUserId(userIdLogin int) ([]kos.RatingCore, error) {
	var kosDataGorm []Rating
	tx := repo.db.Preload("User").Preload("BoardingHouse").Where("user_id = ?", userIdLogin).Find(&kosDataGorm)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var results []kos.RatingCore
	for _, kosData := range kosDataGorm {
		result := kosData.ModelToCoreRating()
		results = append(results, result)
	}
	return results, nil
}
