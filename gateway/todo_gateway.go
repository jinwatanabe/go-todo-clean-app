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
		todo := domain.Todo{
			Id: domain.TodoId{Value: t.Id},
			Title: domain.TodoTitle{Value: t.Title},
			Done: domain.TodoDone{Value: t.Done},
		}
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

	todo := domain.Todo{
		Id:	domain.TodoId{Value: result.Id},
		Title: domain.TodoTitle{Value: result.Title},
		Done: domain.TodoDone{Value: result.Done},
	}

	return todo, nil
}

func (t TodoGateway) Create(todo domain.CreateTodo) (domain.Todo, error) {
	createTodo := driver.CreateTodo{
		Title: todo.Title.Value,
		Done: false,
	}
	
	err := t.todoDriver.Create(createTodo)

	if err != nil {
		return domain.Todo{}, err
	}

	newTodo := domain.NewTodo(
		todo.Title.Value,
		false,
	)

	return newTodo, nil
}

func (t TodoGateway) Update(id domain.TodoId, todo domain.UpdateTodo) (error) {
	intId := id.Value
	updateTodo := driver.UpdateTodo{
		Title: todo.Title.Value,
		Done: todo.Done.Value,
	}

	err := t.todoDriver.Update(intId, updateTodo)
	if err != nil {
		return err
	}

	return nil
}

func (t TodoGateway) Delete(id domain.TodoId) (error) {
	intId := id.Value
	err := t.todoDriver.Delete(intId)

	if err != nil {
		return err
	}

	return nil
}

func ProvideTodoPort(d driver.TodoDriver) port.TodoPort {
	return &TodoGateway{d}
}