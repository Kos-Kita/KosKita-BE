package data

import (
	"KosKita/features/booking"
	kd "KosKita/features/kos/data"
	"KosKita/utils/externalapi"

	"gorm.io/gorm"
)

type bookQuery struct {
	db              *gorm.DB
	paymentMidtrans externalapi.MidtransInterface
}

func New(db *gorm.DB, mid externalapi.MidtransInterface) booking.BookDataInterface {
	return &bookQuery{
		db:              db,
		paymentMidtrans: mid,
	}
}

// Insert implements booking.BookDataInterface.
func (repo *bookQuery) Insert(userIdLogin int, input booking.BookingCore) (*booking.BookingCore, error) {
	boardingHouse := kd.BoardingHouse{}
	if err := repo.db.First(&boardingHouse, input.BoardingHouseId).Error; err != nil {
		return nil, err
	}

	input.Total = float64(boardingHouse.Price)

	bookModel := CoreToModelBook(input)
	bookModel.UserId = uint(userIdLogin)
	bookModel.Total = input.Total
	bookModel.Payment.ExpiredAt = nil
	bookModel.Payment.PaidAt = nil

	if err := bookModel.GenerateCode(); err != nil {
		return nil, err
	}

	if err := repo.db.Create(&bookModel).Error; err != nil {
		return nil, err
	}
	tx := repo.db.Preload("User").Where("user_id = ?", userIdLogin).First(&bookModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	input.Code = bookModel.Code

	payment, errPay := repo.paymentMidtrans.NewOrderPayment(input)

	if errPay != nil {
		return nil, errPay
	}

	bookModel.Payment.Method = payment.Method
	bookModel.Payment.Bank = payment.Bank
	bookModel.Payment.VirtualNumber = payment.VirtualNumber
	bookModel.Payment.BillKey = payment.BillKey
	bookModel.Payment.BillCode = payment.BillCode
	bookModel.Payment.Status = payment.Status
	bookModel.Payment.ExpiredAt = &payment.ExpiredAt
	// bookModel.Payment.PaidAt = &payment.PaidAt
	bookModel.Payment.PaidAt = nil

	if err := repo.db.Save(&bookModel).Error; err != nil {
		return nil, err
	}

	if payment.Status == "settlement" {
		boardingHouse.Rooms -= 1
		if err := repo.db.Save(&boardingHouse).Error; err != nil {
			return nil, err
		}
	}

	bookCore := ModelToCoreBook(bookModel)
	if payment != nil {
		bookCore.Payment = *payment
	}

	return &bookCore, nil
}
