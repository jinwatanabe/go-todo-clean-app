package usecase

import (
	"go-todo-clean-app/domain"

	"github.com/stretchr/testify/mock"
)

type MockTodoPort struct {
	mock.Mock
}

func (m MockTodoPort) GetAll() ([]domain.Todo, error) {
	args := m.Called()
	return args.Get(0).([]domain.Todo), args.Error(1)
}