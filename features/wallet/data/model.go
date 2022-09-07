package data

import (
	"be11/apiclean/features/user"
	"be11/apiclean/features/wallet"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Jenis  string `json:"jenis" form:"jenis"`
	Nomor  string `json:"nomor" form:"nomor"`
	UserID uint   `json:"user_id" form:"user_id"`
	User   User
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
	Wallets  []Wallet
}

func fromCore(dataCore wallet.Core) Wallet {
	dataModel := Wallet{
		Jenis:  dataCore.Jenis,
		Nomor:  dataCore.Nomor,
		UserID: dataCore.UserID,
	}
	return dataModel
}

func (data *Wallet) toCore() wallet.Core {
	return wallet.Core{
		ID:        data.ID,
		Jenis:     data.Jenis,
		Nomor:     data.Nomor,
		UserID:    data.UserID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		User: user.Core{
			ID:      data.User.ID,
			Name:    data.User.Name,
			Email:   data.User.Email,
			Address: data.User.Address,
		},
	}
}

func toCoreList(data []Wallet) []wallet.Core {
	var dataCore []wallet.Core
	for key := range data {
		dataCore = append(dataCore, data[key].toCore())
	}
	return dataCore
}
