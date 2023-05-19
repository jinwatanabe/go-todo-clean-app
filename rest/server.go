package rest

import (
	"go-todo-clean-app/di"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
		r := gin.Default()

		v1 := r.Group("/v1")

		{
			systemHandler := di.InitSystemHandler()
			v1.GET("/systems/ping", systemHandler.Ping)

			todoHandler := di.InitTodoHandler()
			v1.GET("/todos", todoHandler.GetAllTodos)
		}

		return r
}