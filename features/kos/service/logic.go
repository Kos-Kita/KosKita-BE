package service

import (
	"KosKita/features/kos"
	"KosKita/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type kosService struct {
	kosData     kos.KosDataInterface
	userService user.UserServiceInterface
	validate    *validator.Validate
}

func New(repo kos.KosDataInterface, us user.UserServiceInterface) kos.KosServiceInterface {
	return &kosService{
		kosData:     repo,
		validate:    validator.New(),
		userService: us,
	}
}

// Create implements kos.KosServiceInterface.
func (ks *kosService) Create(userIdLogin int, input kos.Core) error {
	user, err := ks.userService.GetById(userIdLogin)
	if err != nil {
		return err
	}

	if user.Role != "owner" {
		return errors.New("lu bukan owner")
	}

	if input.Name == "" {
		return errors.New("lu belum buat name")
	}

	if input.Category == "" {
		return errors.New("category nya mana bos")
	}

	if input.Price <= 0 {
		return errors.New("harga nya lu belum isi")
	}

	if input.Rooms <= 0 {
		return errors.New("isi rooms nya")
	}

	if input.Address == "" {
		return errors.New("alamat lu dimana")
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

// Put implements kos.KosServiceInterface.
func (ks *kosService) Put(userIdLogin int, input kos.Core) error {
	err := ks.kosData.Update(userIdLogin, input)
	if err != nil {
		return err
	}
	return nil
}

// CreateRating implements kos.KosServiceInterface.
func (ks *kosService) CreateRating(userIdLogin int, kosId int, input kos.RatingCore) error {
	if input.Score < 1 || input.Score > 5 {
		return errors.New("skor rating harus antara 1 dan 5")
	}

	errInsert := ks.kosData.InsertRating(userIdLogin, kosId, input)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

// GetByRating implements kos.KosServiceInterface.
func (ks *kosService) GetByRating() ([]kos.RatingCore, error) {
	result, err := ks.kosData.SelectByRating()
	return result, err
}
