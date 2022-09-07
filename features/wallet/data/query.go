package data

import (
	"be11/apiclean/features/wallet"

	"gorm.io/gorm"
)

type walletData struct {
	db *gorm.DB
}

func New(db *gorm.DB) wallet.DataInterface {
	return &walletData{
		db: db,
	}
}

// func (repo *walletData) SelectAllWallet() (wallet.Core, error){
// 	return
// }
