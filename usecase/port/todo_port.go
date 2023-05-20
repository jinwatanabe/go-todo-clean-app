package port

import "go-todo-clean-app/domain"

type TodoPort interface {
	GetAll() ([]domain.Todo, error)
	GetById(id domain.TodoId) (domain.Todo, error)
}