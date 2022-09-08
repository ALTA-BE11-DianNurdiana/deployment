package factory

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	authData "be11/apiclean/features/auth/data"
	authDelivery "be11/apiclean/features/auth/delivery"
	authUsecase "be11/apiclean/features/auth/usecase"

	userData "be11/apiclean/features/user/data"
	userDelivery "be11/apiclean/features/user/delivery"
	userUsecase "be11/apiclean/features/user/usecase"

	walletData "be11/apiclean/features/wallet/data"
	walletDelivery "be11/apiclean/features/wallet/delivery"
	walletUsecase "be11/apiclean/features/wallet/usecase"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authDataFactory := authData.New(db)
	authUsecaseFactory := authUsecase.New(authDataFactory)
	authDelivery.New(e, authUsecaseFactory)

	userDataFactory := userData.New(db)
	userUsecaseFactory := userUsecase.New(userDataFactory)
	userDelivery.New(e, userUsecaseFactory)

	walletDataFactory := walletData.New(db)
	walletUsecaseFactory := walletUsecase.New(walletDataFactory)
	walletDelivery.New(e, walletUsecaseFactory)

}
