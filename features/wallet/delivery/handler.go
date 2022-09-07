package delivery

import (
	"be11/apiclean/features/wallet"
)

type WalletDelivery struct {
	walletUsecase wallet.UsecaseInterface
}

// func New(e *echo.Echo, usecase wallet.UsecaseInterface) {
// 	handler := &WalletDelivery{
// 		walletUsecase: usecase,
// 	}

// 	// e.GET("/wallets", handler.GetAllWallet)
// }

// func (delivery *WalletDelivery) GetAllWallet(c echo.Context) error {
// 	results, err := delivery.walletUsecase.GetAllWallet()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
// 	}

// 	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", fromCoreList(results)))
// }
