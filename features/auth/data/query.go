package data

import (
	"be11/apiclean/features/auth"
	"be11/apiclean/middlewares"
	"errors"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.DataInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) InsertData(data auth.Core) (string, error) {
	var dataLogin Auth
	tx := repo.db.Create(&dataLogin)
	if tx.Error != nil {
		return "", tx.Error
	}

	if tx.RowsAffected != 1 {
		return "", errors.New("login failed")
	}

	token, errToken := middlewares.CreateToken(int(dataLogin.ID))
	if errToken != nil {
		return "", errToken
	}
	return token, nil
}
