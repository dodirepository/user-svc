package usecases

import (
	"github.com/dodirepository/user-svc/internal/adapters/repository"
	repo "github.com/dodirepository/user-svc/internal/domain/repository"
	usecases "github.com/dodirepository/user-svc/internal/domain/usecases"
)

type UserUsecase struct {
	userRepository repo.UserCrud
}

func UserUsecaseHandler() usecases.UserUsecaseInterface {
	return &UserUsecase{
		userRepository: repository.UserRepositoryHandler(),
	}
}
