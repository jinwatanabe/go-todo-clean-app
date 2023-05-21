package usecase

import (
	"go-todo-clean-app/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	todoPort := new(MockTodoPort)
	usecase := TodoUsecase{todoPort}
	todoPort.On(("GetAll")).Return([]domain.Todo{}, nil)
	actual, _ := usecase.GetAll()
	assert.Equal(t, []domain.Todo{}, actual)
}

func Test_GetById(t *testing.T) {
	todoPort := new(MockTodoPort)
	usecase := TodoUsecase{todoPort}
	todoPort.On("GetById", domain.TodoId{Value: 1}).Return(domain.Todo{}, nil)
	actual, _ := usecase.GetById(domain.TodoId{Value: 1})
	assert.Equal(t, domain.Todo{}, actual)
}

func Test_Create(t *testing.T) {
	todoPort := new(MockTodoPort)
	usecase := TodoUsecase{todoPort}
	todoPort.On("Create", domain.CreateTodo{Title: domain.TodoTitle{Value: "title"}}).Return(domain.Todo{}, nil)
	actual, _ := usecase.Create(domain.CreateTodo{Title: domain.TodoTitle{Value: "title"}})
	assert.Equal(t, domain.Todo{}, actual)
}

func Test_Update(t *testing.T) {
	todoPort := new(MockTodoPort)
	usecase := TodoUsecase{todoPort}
	todoPort.On("Update", domain.TodoId{Value: 1}, domain.UpdateTodo{Title: domain.TodoTitle{Value: "title"}}).Return(domain.Todo{}, nil)
	err := usecase.Update(domain.TodoId{Value: 1}, domain.UpdateTodo{Title: domain.TodoTitle{Value: "title"}})
	assert.Nil(t, err)
}