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

func (u TodoUsecase) GetById(id domain.TodoId) (domain.Todo, error) {
	todo, err := u.todoPort.GetById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (u TodoUsecase) Create(todo domain.CreateTodo) (domain.Todo, error) {
	newTodo, err := u.todoPort.Create(todo)

	if err != nil {
		return domain.Todo{}, err
	}

	return newTodo, nil
}

func ProvideTodoUsecase(todoPort port.TodoPort) TodoUsecase {
	return TodoUsecase{todoPort}
}