package gateway

import (
	"go-todo-clean-app/domain"
	"go-todo-clean-app/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {

	mockDriver := new(MockTodoDriver)
	todo := []driver.Todo {driver.NewTodo(1, "title", false)}
	mockDriver.On("GetAll").Return(todo, nil)

	gateway := TodoGateway{mockDriver}
	actual, _ := gateway.GetAll()
	expected := []domain.Todo {domain.NewTodo("title", false)}

	assert.Equal(t, expected, actual)
}

func Test_GetById(t *testing.T) {

	MockDriver := new(MockTodoDriver)
	todo := driver.NewTodo(
		1,
		"title",
		false,
	)
	gateway := TodoGateway{MockDriver}
	MockDriver.On("GetById", 1).Return(todo, nil)

	actual, _ := gateway.GetById(domain.TodoId{Value: 1})
	expected := domain.NewTodo(
		"title",
		false,
	)

	assert.Equal(t, expected, actual)
}

func Test_Create(t *testing.T) {
	MockDriver := new(MockTodoDriver)
	gateway := TodoGateway{MockDriver}
	MockDriver.On("Create", driver.CreateTodo{Title: "title", Done: false}).Return(domain.Todo{}, nil)

	actual, _ := gateway.Create(domain.CreateTodo{Title: domain.TodoTitle{Value: "title"}})
	expected := domain.Todo{Title: domain.TodoTitle{Value: "title"}, Done: domain.TodoDone{Value: false}}

	assert.Equal(t, expected, actual)
}