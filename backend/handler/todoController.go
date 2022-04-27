package handler

import (
	"assignment/model"
	"assignment/service"
	"encoding/json"
	"net/http"
)

type ITodoController interface {
	Handle(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

var todoCtrl = TodoController{}

type TodoController struct {
	srv service.ITodoService
}

func NewTodoController(srv service.ITodoService) ITodoController {
	todoCtrl.srv = srv
	return &todoCtrl
}

func (todoCtrl *TodoController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/api/todo-item" {
			todoCtrl.Create(w, r)
		}
	case http.MethodGet:
		if r.URL.Path == "/api/todo-items" {
			todoCtrl.GetAll(w, r)
		}
	}
}

func (todoCtrl *TodoController) Create(w http.ResponseWriter, r *http.Request) {
	var item = model.TodoItem{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	todoCtrl.srv.Create(&item)

	w.WriteHeader(http.StatusCreated)
}

func (todoCtrl *TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	items := todoCtrl.srv.GetAll()

	marshal, err := json.Marshal(items)
	w.Header().Add("content-type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshal)
}
