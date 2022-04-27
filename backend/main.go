package main

import (
	"assignment/server"
)

func main() {
	serverObj := server.NewServer()
	serverObj.StartServer(3000)
}
