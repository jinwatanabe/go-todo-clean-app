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