package data

import (
	"KosKita/features/user"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.UserDataInterface.
func (repo *userQuery) Insert(input user.Core) error {
	dataGorm := CoreToModel(input)

	tx := repo.db.Create(&dataGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectById implements user.UserDataInterface.
func (repo *userQuery) SelectById(userId int) (*user.Core, error) {
	panic("unimplemented")
}

// Update implements user.UserDataInterface.
func (repo *userQuery) Update(userId int, input user.Core) error {
	panic("unimplemented")
}

// Delete implements user.UserDataInterface.
func (repo *userQuery) Delete(userId int) error {
	panic("unimplemented")
}

// Login implements user.UserDataInterface.
func (repo *userQuery) Login(email string, password string) (data *user.Core, err error) {
	panic("unimplemented")
}
