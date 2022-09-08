package delivery

import (
	"be11/apiclean/features/auth"
	"be11/apiclean/utils/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthDelivery struct {
	authUsecase auth.UsecaseInterface
}

func New(e *echo.Echo, usecase auth.UsecaseInterface) {
	handler := &AuthDelivery{
		authUsecase: usecase,
	}

	e.POST("/login", handler.Login)
}

func (delivery *AuthDelivery) Login(c echo.Context) error {
	var dataRequest AuthRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}

	token, errToken := delivery.authUsecase.Login(toCore(dataRequest))
	if errToken != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed login"))
	}
	dataToken := map[string]interface{}{
		"token": token,
	}
	return c.JSON(http.StatusCreated, helper.SuccessDataResponseHelper("success login", dataToken))
}
