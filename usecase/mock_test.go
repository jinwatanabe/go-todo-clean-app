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

func (m MockTodoPort) GetById(id domain.TodoId) (domain.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m MockTodoPort) Create(todo domain.CreateTodo) (domain.Todo, error) {
	args := m.Called(todo)
	return args.Get(0).(domain.Todo), args.Error(1)
}

func (m MockTodoPort) Update(id domain.TodoId, todo domain.UpdateTodo) (error) {
	args := m.Called(id, todo)
	return args.Error(1)
}