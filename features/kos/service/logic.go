package service

import (
	"KosKita/features/kos"
	"KosKita/features/user"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type kosService struct {
	kosData  kos.KosDataInterface
	userService user.UserServiceInterface 
	validate *validator.Validate
}

func New(repo kos.KosDataInterface) kos.KosServiceInterface {
	return &kosService{
		kosData:  repo,
		validate: validator.New(),
	}
}

// Create implements kos.KosServiceInterface.
func (ks *kosService) Create(userIdLogin int, input kos.Core) error {
	user, err :=  ks.userService.GetById(userIdLogin)
	if err != nil {
		return err
	}
	
	fmt.Println("user - ", user.Role)

	if user.Role != "owner" {
		return errors.New("only owners can create kos")
	}

	if input.Name == "" {
		return errors.New("name is required")
	}

	if input.Category == "" {
		return errors.New("category is required")
	}

	if input.Price <= 0 {
		return errors.New("price must be positive")
	}

	if input.Rooms <= 0 {
		return errors.New("rooms must be positive")
	}

	if input.Address == "" {
		return errors.New("address is required")
	}

	if input.PhotoMain == "" {
		return errors.New("photo main is required")
	}

	if input.PhotoFront == "" {
		return errors.New("photo front is required")
	}

	if input.PhotoBack == "" {
		return errors.New("photo back is required")
	}

	if input.PhotoRoomFront == "" {
		return errors.New("photo room front is required")
	}

	if input.PhotoRoomInside == "" {
		return errors.New("photo room inside is required")
	}
	
	errValidate := ks.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	errInsert := ks.kosData.Insert(userIdLogin, input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}
