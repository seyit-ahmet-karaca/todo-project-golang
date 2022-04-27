package model

type TodoItem struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type TodoItems []*TodoItem