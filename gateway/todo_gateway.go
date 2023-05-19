package gateway

import (
	"go-todo-clean-app/domain"
	"go-todo-clean-app/driver"
	"go-todo-clean-app/usecase/port"
)

type TodoGateway struct {
	todoDriver driver.TodoDriver
}

func (t TodoGateway) GetAll() ([]domain.Todo, error) {
	result, err := t.todoDriver.GetAll()

	if err != nil {
		return nil, err
	}

	var todos []domain.Todo

	for _, t := range result {
		todo := domain.NewTodo(t.Title, t.Done)
		todos = append(todos, todo)
	}

	return todos, nil

}

func ProvideTodoPort(d driver.TodoDriver) port.TodoPort {
	return &TodoGateway{d}
}