package service

import (
	"KosKita/features/kos"
	"KosKita/features/user"
	"KosKita/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	returnData := kos.CoreInput{
		UserID:      1,
		Name:        "KOS",
		Description: "kos ini murah sekali",
		Category:    "putra",
		Price:       100000,
		Rooms:       10,
		Address:     "Jl. Kost Murah No.1",
		Longitude:   "123.456",
		Latitude:    "98.768",
		KosFacilities: []kos.KosFacilityCore{
			{
				Facility: "AC",
			},
			{
				Facility: "TV",
			},
		},
		KosRules: []kos.KosRuleCore{
			{
				Rule: "Dilarang Merokok",
			},
			{
				Rule: "Dilarang Membawa Hewan Peliharaan",
			},
		},
	}

	userIdLogin := int(returnData.UserID)

	t.Run("user is not owner", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "renter"}, nil).Once()

		kosId, err := srv.Create(userIdLogin, returnData)

		assert.Error(t, err)
		assert.Equal(t, uint(0), kosId)
		assert.EqualError(t, err, "anda bukan owner")
	})

	t.Run("name is empty", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		returnData.Name = ""
		_, err := srv.Create(userIdLogin, returnData)

		assert.Error(t, err)
		assert.EqualError(t, err, "name anda kosong")
		returnData.Name = "Kos Murah"
	})

	t.Run("category is empty", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		returnData.Category = ""
		_, err := srv.Create(userIdLogin, returnData)

		assert.Error(t, err)
		assert.EqualError(t, err, "category anda kosong")
		returnData.Category = "putra"
	})

	t.Run("price is zero or less", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		returnData.Price = 0
		_, err := srv.Create(userIdLogin, returnData)

		assert.Error(t, err)
		assert.EqualError(t, err, "harga anda kosong")
		returnData.Price = 1000000
	})

	t.Run("rooms is zero or less", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		returnData.Rooms = 0
		_, err := srv.Create(userIdLogin, returnData)

		assert.Error(t, err)
		assert.EqualError(t, err, "rooms anda kosong")
		returnData.Rooms = 10
	})

	t.Run("address is empty", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		returnData.Address = ""
		_, err := srv.Create(userIdLogin, returnData)

		assert.Error(t, err)
		assert.EqualError(t, err, "alamat anda kosong")
		returnData.Address = "Jl. Kost Murah No.1"
	})
	
	t.Run("error from repository", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil)
		repo.On("Insert", userIdLogin, returnData).Return(uint(0), errors.New("database error")).Once()
		
		kosId, err := srv.Create(userIdLogin, returnData)
		
		assert.Error(t, err)
		assert.NotNil(t, kosId)
		assert.Equal(t, uint(0), kosId)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})
	
	t.Run("success", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil)
		repo.On("Insert", userIdLogin, returnData).Return(returnData.ID, nil).Once()

		kosId, err := srv.Create(userIdLogin, returnData)

		assert.NoError(t, err)
		assert.Equal(t, returnData.ID, kosId)

		repo.AssertExpectations(t)
	})
}