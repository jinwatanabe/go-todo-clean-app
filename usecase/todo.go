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

func (u TodoUsecase) Update(id domain.TodoId, todo domain.UpdateTodo) (error) {
	err := u.todoPort.Update(id, todo)

	if err != nil {
		return err
	}

	return nil
}

func (u TodoUsecase) Delete(id domain.TodoId) (error) {
	err := u.todoPort.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func ProvideTodoUsecase(todoPort port.TodoPort) TodoUsecase {
	return TodoUsecase{todoPort}
}