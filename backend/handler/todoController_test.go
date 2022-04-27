// +build unit

package handler

import (
	"assignment/mock"
	"assignment/model"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("Create request with correct json using handler method", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenTitle := "test"
		giventem := &model.TodoItem{Title: givenTitle}
		serviceReturnValue := 1

		mockService := mock.NewMockITodoService(mockController)
		mockService.
			EXPECT().
			Create(giventem).
			Return(serviceReturnValue, nil).
			Times(1)

		controller := NewTodoController(mockService)
		r := httptest.NewRequest(http.MethodPost, "/api/todo-item", bytes.NewBuffer([]byte(`{"title": "`+givenTitle+`"}`)))
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
		assert.Nil(t, w.Body.Bytes())
	})

	t.Run("Create request with invalid data", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		giventem := &model.TodoItem{Title: "test"}
		serviceReturnValue := 1

		mockService := mock.NewMockITodoService(mockController)
		mockService.
			EXPECT().
			Create(giventem).
			Return(serviceReturnValue, nil).
			Times(0)

		controller := NewTodoController(mockService)
		r := httptest.NewRequest(http.MethodPost, "/api/todo-item", bytes.NewBuffer([]byte("some invalid json")))
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
		assert.True(t, w.Body.Bytes() != nil)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("GetAll request using handle method", func(t *testing.T) {
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

		mockService := mock.NewMockITodoService(mockController)
		mockService.
			EXPECT().
			GetAll().
			Return(expectedItems).
			Times(1)

		controller := NewTodoController(mockService)
		r := httptest.NewRequest(http.MethodGet, "/api/todo-items", http.NoBody)
		w := httptest.NewRecorder()

		controller.Handle(w, r)

		var actualItems = model.TodoItems{}
		err := json.NewDecoder(w.Body).Decode(&actualItems)

		assert.Nil(t, err)
		assert.Equal(t, "application/json; charset=UTF-8", w.Header().Get("Content-Type"))
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Equal(t, expectedItems, actualItems)
	})
}
