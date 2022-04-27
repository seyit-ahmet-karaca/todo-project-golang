package repository

import (
	"assignment/data"
	"assignment/model"
)

type ITodoRepository interface {
	Insert(item *model.TodoItem) *model.TodoItem
	FindAll() model.TodoItems
}

type TodoRepository struct{}

func NewTodoRepository() ITodoRepository{
	return &TodoRepository{}
}

func (t *TodoRepository) Insert(item *model.TodoItem) *model.TodoItem {
	indexCreated := len(data.Data)
	item.Id = indexCreated + 1
	data.Data = append(data.Data, item)
	return item
}

func (t *TodoRepository) FindAll() model.TodoItems {
	return data.Data
}
