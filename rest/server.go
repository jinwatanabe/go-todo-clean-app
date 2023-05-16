package rest

import (
	"go-todo-clean-app/di"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
		r := gin.Default()

		systemHandler := di.InitSystemHandler()
		r.GET("/v1/systems/ping", systemHandler.Ping)
		
		return r
}