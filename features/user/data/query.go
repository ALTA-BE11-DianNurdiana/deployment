package data

import (
	"be11/apiclean/features/user"
	"errors"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userData{
		db: db,
	}
}

func (repo *userData) SelectAllData() ([]user.Core, error) {
	var data []User
	tx := repo.db.Find(&data)
	// tx := config.DB.Preload("Wallets").Find(&data)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataCore := toCoreList(data)
	return dataCore, nil
}

func (repo *userData) InsertData(data user.Core) (int, error) {
	dataModel := fromCore(data)
	tx := repo.db.Create(&dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}

	return int(tx.RowsAffected), nil
}

func (repo *userData) SelectById(id int) (user.Core, error) {
	var data User

	tx := repo.db.First(&data, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	dataCore := toCore(data)
	return dataCore, nil
}

func toCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
		Address:  data.Address,
	}
}

func (repo *userData) UpdateData(id int, newData user.Core) (int, error) {
	dataModel := fromCore(newData)
	tx := repo.db.Model(&User{}).Where("id = ?", id).Updates(dataModel)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to update data")
	}

	return 1, nil
}

func (repo *userData) DeleteData(id int) (int, error) {
	var data User
	tx := repo.db.Where("id = ?", id).Delete(&data)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("failed to delete data")
	}

	return int(tx.RowsAffected), nil
}
