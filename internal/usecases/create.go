package usecases

import (
	"net/http"

	pkg "github.com/dodirepository/common-lib"
	domain "github.com/dodirepository/user-svc/internal/domain/repository"
	usecases "github.com/dodirepository/user-svc/internal/domain/usecases"
	"github.com/sirupsen/logrus"
)

func (u *UserUsecase) Insert(req usecases.UserCreate) *usecases.ErrorResponse {
	passHash, err := pkg.HashPassword(req.Password)
	if err != nil {
		logrus.WithError(err).Error("Error hashing password: " + req.Password)
		return &usecases.ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  http.StatusInternalServerError,
		}
	}
	dataCreate := domain.Users{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: passHash,
	}
	err = u.userRepository.Create(dataCreate)
	if err != nil {
		if pkg.IsDuplicateEntryError(err) {
			return &usecases.ErrorResponse{
				Message: "email or phone already exists",
				Status:  http.StatusInternalServerError,
			}
		}
		logrus.WithError(err).Error("Error creating user")
		return &usecases.ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil

}
