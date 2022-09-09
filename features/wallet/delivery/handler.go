package delivery

import (
	"be11/apiclean/features/wallet"
	"be11/apiclean/middlewares"
	"be11/apiclean/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletDelivery struct {
	walletUsecase wallet.UsecaseInterface
}

func New(e *echo.Echo, usecase wallet.UsecaseInterface) {
	handler := &WalletDelivery{
		walletUsecase: usecase,
	}

	e.GET("/wallets", handler.GetAllWallet)
	e.POST("/wallets", handler.PostWallet, middlewares.JWTMiddleware())
	// e.GET("/wallets/:id", handler.GetAllWallet)
}

func (delivery *WalletDelivery) GetAllWallet(c echo.Context) error {
	results, err := delivery.walletUsecase.GetAllWallet()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", fromCoreList(results)))
}

func (delivery *WalletDelivery) PostWallet(c echo.Context) error {

	var dataRequest WalletRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}

	idToken := middlewares.ExtractToken(c)

	dataCore := toCore(dataRequest)
	dataCore.UserID = uint(idToken)
	row, err := delivery.walletUsecase.PostWallet(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}

// Get wallet by id
// func (delivery *WalletDelivery) GetById(c echo.Context) error {
// 	id := c.Param("id")
// 	idConv, errConv := strconv.Atoi(id)

// 	result, errConv := delivery.walletUsecase.GetById(idConv)
// 	if errConv != nil {
// 		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
// 	}

// 	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", fromCore(result)))
// }
