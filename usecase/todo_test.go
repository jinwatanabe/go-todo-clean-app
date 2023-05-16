package usecase

import (
	"go-todo-clean-app/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAllTodo_Execxute(t *testing.T) {
	todoPort := new(MockTodoPort)
	usecase := TodoUsecase{todoPort}
	todoPort.On(("GetAll")).Return([]domain.Todo{}, nil)
	actual, _ := usecase.GetAll()
	assert.Equal(t, []domain.Todo{}, actual)
}