package user

import "time"

type Core struct {
	ID           uint
	Name         string 
	UserName     string 
	Email        string 
	Password     string 
	Gender       string 
	Role         string
	PhotoProfile string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input Core) error
	SelectById(userId int) (*Core, error)
	Update(userId int, input Core) error
	Delete(userId int) error
	Login(email, password string) (data *Core, err error)
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input Core) error
	GetById(userId int) (*Core, error)
	Update(userId int, input Core) error
	Delete(userId int) error
	Login(email, password string) (data *Core, token string, err error)
}
