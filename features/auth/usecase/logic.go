package usecase

import (
	"be11/apiclean/features/auth"
)

type authUsecase struct {
	authData auth.DataInterface
}

func New(data auth.DataInterface) auth.UsecaseInterface {
	return &authUsecase{
		authData: data,
	}
}

func (usecase *authUsecase) Login(data auth.Core) (string, error) {
	result, err := usecase.authData.InsertData(data)
	if err != nil {
		return "", err
	}
	return result, nil
}
