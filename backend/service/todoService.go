package service

import (
	"assignment/model"
	"assignment/repository"
	"errors"
)

type ITodoService interface {
	Create(*model.TodoItem) (int, error)
	GetAll() model.TodoItems
}

type TodoService struct {
	repo repository.ITodoRepository
}

var srv = TodoService{}

func NewTodoService(repo repository.ITodoRepository) ITodoService {
	srv.repo = repo
	return &srv
}

func (t *TodoService) Create(item *model.TodoItem) (int, error) {
	if item == nil {
		return 0, errors.New("Item cannot be nil")
	}

	insertedItem := t.repo.Insert(item)
	return insertedItem.Id, nil
}

func (t *TodoService) GetAll() model.TodoItems {
	return t.repo.FindAll()
}
