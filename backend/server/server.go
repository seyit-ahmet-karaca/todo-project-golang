package server

import (
	"assignment/config"
	"assignment/handler"
	"assignment/repository"
	"assignment/service"
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) {
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository)
	controller := handler.NewTodoController(todoService)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{config.Get().UIUrl},
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", port),
		c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controller.Handle(w, r)
	})))

	if err != nil {
		panic(err)
	}
}
