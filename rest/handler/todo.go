package handler

import (
	"fmt"
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

	fmt.Println(todos)

	response := TodosResponse{
		Todos: todos,
	}

	c.JSON(http.StatusOK, response)
}

func ProvideTodoHandler(u usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{u}
}

type TodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}
