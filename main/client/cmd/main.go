package main

import (
	"github.com/kolodach/golang-todo/cmd"
	"github.com/kolodach/golang-todo/inmemory"
)

func main() {
	serv := inmemory.NewTodoService()
	c := cmd.CreateClient(serv)
	c.Process()
}
