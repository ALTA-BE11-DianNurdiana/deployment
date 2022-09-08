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

func (repo *walletData) SelectAllWallet() ([]wallet.Core, error) {
	var data []Wallet
	tx := repo.db.Preload("Wallet").Find(&data)
	// tx := config.DB.Preload("Wallets").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(data)
	return dataCore, nil
}

func (repo *walletData) InsertWallet(data wallet.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}
