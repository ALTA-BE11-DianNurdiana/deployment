package usecase

import (
	"be11/apiclean/features/user"
	"be11/apiclean/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []user.Core{{Name: "coco", Email: "coco@gmail.id", Password: "123"}}

	t.Run("Succes get all data", func(t *testing.T) {
		repo.On("SelectAllData", mock.Anything).Return(returnData, nil).Once()

		usecase := New(repo)

		res, err := usecase.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, returnData[0].ID, res[0].ID)
		repo.AssertExpectations(t)
	})

	// t.Run("Error get all data", func(t *testing.T) {
	// 	repo.On("SelectAllData", "").Return(nil, errors.New("data not found")).Once()

	// 	srv := New(repo)

	// 	res, err := srv.GetAll()
	// 	assert.Error(t, err)
	// 	assert.Nil(t, res)
	// 	repo.AssertExpectations(t)
	// })
}

func TestPostData(t *testing.T) {
	repo := new(mocks.UserData)
	insertData := user.Core{Name: "coco", Email: "coco@gmail.id", Password: "123"}

	t.Run("Success insert", func(t *testing.T) {
		repo.On("InsertData", mock.Anything).Return(1, nil).Once()

		usecase := New(repo)
		res, err := usecase.PostData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed no name", func(t *testing.T) {
		// repo.On("InsertData", mock.Anything).Return(0, errors.New("there is some error")).Once()
		insertData := user.Core{Email: "coco@gmail.id", Password: "123"}
		usecase := New(repo)
		res, err := usecase.PostData(insertData)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		repo.AssertExpectations(t)
	})
}

func TestSelectById(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := user.Core{Name: "coco", Email: "coco@gmail.id", Password: "123"}

	t.Run("success get data", func(t *testing.T) {
		repo.On("SelectById", mock.Anything).Return(returnData, nil).Once()

		usecase := New(repo)
		res, error := usecase.GetById(1)
		assert.NoError(t, error)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})
}
