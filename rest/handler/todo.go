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

func (h TodoHandler) Create(c *gin.Context) {
	title := c.PostForm("title")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title is required",
		})
		return
	}

	todo := domain.CreateTodo{
		Title: domain.TodoTitle{Value: title},
	}

	newTodo, err := h.todoUsecase.Create(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := TodoResponse{
		Todo: newTodo,
	}

	c.JSON(http.StatusOK, response)
}

func (h TodoHandler) Update(c *gin.Context) {
	paramsId := c.Params.ByName("id")
	intId, err := strconv.Atoi(paramsId)
	id := domain.TodoId{Value: intId}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	title := c.PostForm("title")
	d := c.PostForm("done")

	if title == "" && d == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title or done is required",
		})
		return
	}
	
	done, err := strconv.ParseBool(d)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	updateTodo := domain.UpdateTodo{
		Title: domain.TodoTitle{Value: title},
		Done: domain.TodoDone{Value: done},
	}

	err = h.todoUsecase.Update(id, updateTodo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h TodoHandler) Delete(c *gin.Context) {
	paramsId := c.Params.ByName("id")
	Intid, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message": err.Error(),
		})

		return
	}

	id := domain.TodoId{Value: Intid}
	err = h.todoUsecase.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H {
		"message": "success",
	})
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