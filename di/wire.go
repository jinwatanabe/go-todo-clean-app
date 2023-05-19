//go:build wireinject
// +build wireinject

package di

import (
	"go-todo-clean-app/rest/handler"
	"go-todo-clean-app/driver"
	"go-todo-clean-app/gateway"
	"go-todo-clean-app/usecase"

	"github.com/google/wire"
)

func InitSystemHandler() *handler.SystemHandler {
	wire.Build(handler.NewSystemHandler)
	return nil
}

func InitTodoHandler() *handler.TodoHandler {
	wire.Build(
		handler.ProvideTodoHandler,
		usecase.ProvideTodoUsecase,
		gateway.ProvideTodoPort,
		driver.ProvideTodoDriver,
		driver.ProvideDatabaseConnection,
	)
	return nil
}
