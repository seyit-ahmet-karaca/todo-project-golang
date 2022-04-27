// +build unit

package service

import (
	"assignment/mock"
	"assignment/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("Create TodoItem", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenItemTitle := "test"
		givenItem := model.TodoItem{
			Title: givenItemTitle,
		}

		expectedItem := givenItem
		expectedItem.Id = 1

		mockRepository := mock.NewMockITodoRepository(mockController)
		mockRepository.
			EXPECT().
			Insert(&givenItem).
			Return(&expectedItem).
			Times(1)

		service := NewTodoService(mockRepository)
		actualResult, err := service.Create(&givenItem)

		assert.Equal(t, expectedItem.Id, actualResult)
		assert.Nil(t, err)
	})

	t.Run("Create TodoItem error when given input is nil", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenItem := model.TodoItem{
			Title: "test",
		}

		mockRepository := mock.NewMockITodoRepository(mockController)
		mockRepository.
			EXPECT().
			Insert(nil).
			Return(&givenItem).
			Times(0)

		service := NewTodoService(mockRepository)
		_, err := service.Create(nil)

		assert.True(t, err != nil)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Getall", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		expectedItems := model.TodoItems{
			&model.TodoItem{
				Id:    1,
				Title: "test1",
			},
			&model.TodoItem{
				Id:    2,
				Title: "test2",
			},
		}

		mockRepository := mock.NewMockITodoRepository(mockController)
		mockRepository.
			EXPECT().
			FindAll().
			Return(expectedItems).
			Times(1)

		service := NewTodoService(mockRepository)
		actualItems := service.GetAll()

		assert.Equal(t, expectedItems, actualItems)
	})
}
