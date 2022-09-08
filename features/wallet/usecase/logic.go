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
	if err != nil {
		return []wallet.Core{}, err
	}
	return results, nil
}

func (usecase *walletUsecase) PostWallet(data wallet.Core) (int, error) {
	row, err := usecase.walletData.InsertWallet(data)
	if err != nil {
		return 0, err
	}
	return row, nil
}

// func (usecase *walletUsecase) GetById(id int) (wallet.Core, error) {
// 	result, err := usecase.walletData.SelectById(id)
// 	if err != nil {
// 		return wallet.Core{}, err
// 	}
// 	return result, nil
// }
