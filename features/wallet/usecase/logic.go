package usecase

import (
	"be11/apiclean/features/wallet"
)

type walletUsecase struct {
	walletData wallet.DataInterface
}

func New(data wallet.DataInterface) wallet.UsecaseInterface {
	return &walletUsecase{
		walletData: data,
	}
}

func (usecase *walletUsecase) GetAllWallet() ([]wallet.Core, error) {
	results, err := usecase.walletData.SelectAllWallet()
	return results, err
}

func (usecase *walletUsecase) PostWallet(data wallet.Core) (int, error) {
	row, err := usecase.walletData.InsertWallet(data)
	return row, err
}
