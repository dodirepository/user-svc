package usecases

import (
	"net/http"

	usecases "github.com/dodirepository/user-svc/internal/domain/usecases"
	"github.com/dodirepository/user-svc/internal/middleware"
	"github.com/dodirepository/user-svc/pkg"
	"github.com/sirupsen/logrus"
)

func (u *UserUsecase) Login(userLogin usecases.UserLogin) (interface{}, *usecases.ErrorResponse) {
	users, err := u.userRepository.GetByPhoneOrEmail(userLogin.Username)
	if err != nil {
		return nil, &usecases.ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  http.StatusInternalServerError,
		}
	}
	if users == nil {
		logrus.Info("user not found")
		return nil, &usecases.ErrorResponse{
			Message: "Invalid login credentials",
			Status:  http.StatusUnauthorized,
		}
	}

	if ok := pkg.CheckPasswordHash(userLogin.Password, users.Password); !ok {
		return nil, &usecases.ErrorResponse{
			Message: "Invalid login credentials",
			Status:  http.StatusUnauthorized,
		}
	}

	token, err := middleware.GenerateToken(users.ID, users.Email, users.Phone)
	if err != nil {
		return nil, &usecases.ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  http.StatusInternalServerError,
		}
	}
	result := usecases.UserLoginResponse{
		AccessToken: token,
	}

	return result, nil

}
