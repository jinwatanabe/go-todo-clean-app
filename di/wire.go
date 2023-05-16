//go:build wireinject
// +build wireinject

package di

import (
	"go-todo-clean-app/rest/handler"

	"github.com/google/wire"
)

func InitSystemHandler() *handler.SystemHandler {
	wire.Build(handler.NewSystemHandler)
	return nil
}
