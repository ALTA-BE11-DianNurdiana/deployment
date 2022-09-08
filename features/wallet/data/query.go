package data

import (
	"be11/apiclean/features/wallet"
	"be11/apiclean/utils/helper"
	"errors"

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

func (repo *walletData) SelectAllWallet() ([]wallet.Core, error) {
	var data []Wallet
	tx := repo.db.Find(&data)
	// tx := config.DB.Preload("Wallets").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(data)
	return dataCore, nil
}

func (repo *walletData) InsertWallet(data wallet.Core) (int, error) {
	dataModel := FromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrEmptySlice) {
			return 0, helper.CheckQueryError(tx.Error)
		}
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

// func (repo *walletData) SelectById(id int) (wallet.Core, error) {
// 	var data Wallet

// 	tx := repo.db.First(&data, id)
// 	if tx.Error != nil {
// 		return wallet.Core{}, tx.Error
// 	}

// 	dataCore := toCore(data)
// 	return dataCore, nil
// }

// func (data *Wallet) toCore() wallet.Core {
// 	return wallet.Core{
// 		ID:        data.ID,
// 		Jenis:     data.Jenis,
// 		Nomor:     data.Nomor,
// 		UserID:    data.UserID,
// 		CreatedAt: data.CreatedAt,
// 		UpdatedAt: data.UpdatedAt,
// 		User: user.Core{
// 			ID:      data.User.ID,
// 			Name:    data.User.Name,
// 			Email:   data.User.Email,
// 			Address: data.User.Address,
// 		},
// 	}
// }
