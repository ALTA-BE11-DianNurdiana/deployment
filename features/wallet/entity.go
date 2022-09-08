package wallet

import (
	"be11/apiclean/features/user"
	"time"
)

type Core struct {
	ID        uint
	Jenis     string
	Nomor     string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.Core
}

type UsecaseInterface interface {
	GetAllWallet() (data []Core, err error)
	PostWallet(data Core) (row int, err error)
	// GetById(id int) (data Core, err error)
}

type DataInterface interface {
	SelectAllWallet() (data []Core, err error)
	InsertWallet(data Core) (row int, err error)
	// SelectById(id int) (data Core, err error)
}
