package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SystemHandler struct {
}

func (h SystemHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}
