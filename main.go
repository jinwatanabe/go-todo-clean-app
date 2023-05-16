package main

import (
	"go-todo-clean-app/config"
	"go-todo-clean-app/rest"
)

func main() {
	config.InitConfig()
	s := rest.NewServer()
    s.Run()
}
