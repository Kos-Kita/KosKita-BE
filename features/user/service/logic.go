package service

import (
	"KosKita/features/user"
	"KosKita/utils/encrypts"
	"errors"

	"github.com/go-playground/validator"
)

type userService struct {
	userData    user.UserDataInterface
	hashService encrypts.HashInterface
	validate    *validator.Validate
}

// dependency injection
func New(repo user.UserDataInterface, hash encrypts.HashInterface) user.UserServiceInterface {
	return &userService{
		userData:    repo,
		hashService: hash,
		validate:    validator.New(),
	}
}

// Create implements user.UserServiceInterface.
func (service *userService) Create(input user.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		hashedPass, errHash := service.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errors.New("error hash password")
		}
		input.Password = hashedPass
	}

	err := service.userData.Insert(input)
	return err
}

// GetById implements user.UserServiceInterface.
func (service *userService) GetById(userId int) (*user.Core, error) {
	panic("unimplemented")
}

// Update implements user.UserServiceInterface.
func (service *userService) Update(userId int, input user.Core) error {
	panic("unimplemented")
}

// Delete implements user.UserServiceInterface.
func (service *userService) Delete(userId int) error {
	panic("unimplemented")
}

// Login implements user.UserServiceInterface.
func (service *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	panic("unimplemented")
}
