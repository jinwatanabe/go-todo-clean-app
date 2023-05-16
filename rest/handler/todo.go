package handler

import (
	"go-todo-clean-app/domain"
	"go-todo-clean-app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func (h TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := TodosResponse{
		Todos: todos,
	}

	c.JSON(http.StatusOK, response)
}

type TodosResponse struct {
	Todos []domain.Todo `json:"todo"`
}
