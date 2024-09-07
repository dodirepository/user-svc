package router

import (
	"github.com/dodirepository/user-svc/internal/adapters/controllers/handlers"
	"github.com/dodirepository/user-svc/internal/usecases"
)

var healthCheckControler = handlers.Healthcek{}

var userController = &handlers.UsersHandlersController{
	Usecase: usecases.UserUsecaseHandler(),
}
