package data

import (
	"KosKita/features/user"

	"gorm.io/gorm"
)

// struct user gorm model
type User struct {
	gorm.Model
	Name         string
	UserName     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	Gender       string `gorm:"not null"`
	Role         string `gorm:"not null"`
	PhotoProfile string
}

func CoreToModel(input user.Core) User {
	return User{
		Name:         input.Name,
		UserName:     input.UserName,
		Email:        input.Email,
		Password:     input.Password,
		Gender:       input.Gender,
		Role:         input.Role,
		PhotoProfile: input.PhotoProfile,
	}
}
