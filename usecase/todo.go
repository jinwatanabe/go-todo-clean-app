package usecase

import (
	"go-todo-clean-app/domain"
	"go-todo-clean-app/usecase/port"
)

type TodoUsecase struct {
	todoPort port.TodoPort
}

func (u TodoUsecase) GetAll() ([]domain.Todo, error) {
	return u.todoPort.GetAll()
}