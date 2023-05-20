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

func (t TodoGateway) GetById(id domain.TodoId) (domain.Todo, error) {
	intId := id.Value
	result, err := t.todoDriver.GetById(intId)

	if err != nil {
		return domain.Todo{}, err
	}

	todo := domain.NewTodo(
		result.Title,
		result.Done,
	)

	return todo, nil
}

func ProvideTodoPort(d driver.TodoDriver) port.TodoPort {
	return &TodoGateway{d}
}