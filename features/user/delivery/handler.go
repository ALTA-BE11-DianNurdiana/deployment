package delivery

import (
	"be11/apiclean/features/user"
	"be11/apiclean/utils/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUsecase user.UsecaseInterface
}

func New(e *echo.Echo, usecase user.UsecaseInterface) {
	handler := &UserDelivery{
		userUsecase: usecase,
	}

	e.GET("/users", handler.GetAll)
	e.POST("/users", handler.PostData)
	e.GET("/users/:id", handler.GetById)
	e.PUT("/users/:id", handler.PutData)
	e.DELETE("/users/:id", handler.DeleteData)
}

func (delivery *UserDelivery) GetAll(c echo.Context) error {
	results, err := delivery.userUsecase.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", fromCoreList(results)))
}

func (delivery *UserDelivery) PostData(c echo.Context) error {
	var dataRequest UserRequest
	errBind := c.Bind(&dataRequest)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponseHelper("error bind data"))
	}
	row, err := delivery.userUsecase.PostData(toCore(dataRequest))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error insert data"))
	}

	return c.JSON(http.StatusCreated, helper.SuccessResponseHelper("success insert data"))
}

func (delivery *UserDelivery) GetById(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)

	result, errConv := delivery.userUsecase.GetById(idConv)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error get data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessDataResponseHelper("success", fromCore(result)))
}

func (delivery *UserDelivery) PutData(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("id must be number"))
	}

	var dataUpdate UserRequest
	errBind := c.Bind(&dataUpdate)
	if errBind != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("error bind data"))
	}

	row, err := delivery.userUsecase.PutData(idConv, toCore(dataUpdate))
	if err != nil || row < 1 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed update data"))
	}

	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success update data"))
}

func (delivery *UserDelivery) DeleteData(c echo.Context) error {
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("id must be number"))
	}

	row, err := delivery.userUsecase.DeleteData(idConv)
	if err != nil || row == 0 {
		return c.JSON(http.StatusInternalServerError, helper.FailedResponseHelper("failed delete data"))
	}
	return c.JSON(http.StatusOK, helper.SuccessResponseHelper("success delete data"))
}
