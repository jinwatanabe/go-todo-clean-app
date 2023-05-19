package usecase

import (
	"go-todo-clean-app/domain"
	"go-todo-clean-app/usecase/port"
)

type TodoUsecase struct {
	todoPort port.TodoPort
}

func (u TodoUsecase) GetAll() ([]domain.Todo, error) {
	todos, err :=  u.todoPort.GetAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func ProvideTodoUsecase(todoPort port.TodoPort) TodoUsecase {
	return TodoUsecase{todoPort}
}