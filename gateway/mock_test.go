package gateway

import (
	"go-todo-clean-app/driver"

	"github.com/stretchr/testify/mock"
)

type MockTodoDriver struct {
	mock.Mock
}

func (m MockTodoDriver) GetAll () ([]driver.Todo, error) {
	args := m.Called()
	return args.Get(0).([]driver.Todo), args.Error(1)
}

func (m MockTodoDriver) GetById (id int) (driver.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(driver.Todo), args.Error(1)
}

func (m MockTodoDriver) Create (todo driver.CreateTodo) (error) {
	args := m.Called(todo)
	return args.Error(1)
}

func (m MockTodoDriver) Update (id int, todo driver.UpdateTodo) (error) {
	args := m.Called(id, todo)
	return args.Error(0)
}

func (m MockTodoDriver) Delete (id int) (error) {
	args := m.Called(id)
	return args.Error(0)
}