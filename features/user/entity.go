package user

import (
	"time"
)

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Wallets   []WalletCore
}

type WalletCore struct {
	ID        uint
	Jenis     string
	Nomor     string
	CreatedAt time.Time
}

type UsecaseInterface interface {
	GetAll() (data []Core, err error)
	PostData(data Core) (row int, err error)
	GetById(id int) (data Core, err error)
	PutData(id int, newData Core) (row int, err error)
	DeleteData(id int) (row int, err error)
}

type DataInterface interface {
	SelectAllData() (data []Core, err error)
	InsertData(data Core) (row int, err error)
	SelectById(id int) (data Core, err error)
	UpdateData(id int, newData Core) (row int, err error)
	DeleteData(id int) (row int, err error)
}
