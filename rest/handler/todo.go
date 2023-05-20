package handler

import (
	"go-todo-clean-app/domain"
	"go-todo-clean-app/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func (h TodoHandler) GetAll(c *gin.Context) {
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

func (h TodoHandler) GetById(c *gin.Context) {
	paramsId := c.Params.ByName("id")
	intId, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	id := domain.TodoId{Value: intId}
	todo, err := h.todoUsecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := TodoResponse{
		Todo: todo,
	}

	c.JSON(http.StatusOK, response)

}

func ProvideTodoHandler(u usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{u}
}

type TodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type TodoResponse struct {
	Todo domain.Todo `json:"todo"`
}