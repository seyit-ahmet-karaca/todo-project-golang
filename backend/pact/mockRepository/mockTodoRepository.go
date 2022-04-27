package mockRepository

import (
	"assignment/model"
	"assignment/repository"
)

type MockTodoRepository struct {

}

func NewMockTodoRepository() repository.ITodoRepository{
	return &MockTodoRepository{}
}

func (m *MockTodoRepository) Insert(item *model.TodoItem) *model.TodoItem {
	return &model.TodoItem{
		Id:    1,
		Title: "testItem",
	}
}

func (m *MockTodoRepository) FindAll() model.TodoItems {
	return model.TodoItems{
		&model.TodoItem{
			Id:    1,
			Title: "testItem",
		},
	}
}
