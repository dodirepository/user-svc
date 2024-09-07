package usecases

import (
	"net/http"

	usecases "github.com/dodirepository/user-svc/internal/domain/usecases"
)

func (u *UserUsecase) Detail(ID int64) (interface{}, *usecases.ErrorResponse) {
	users, err := u.userRepository.GetByID(ID)
	if err != nil {
		return nil, &usecases.ErrorResponse{
			Message: http.StatusText(http.StatusInternalServerError),
			Status:  http.StatusInternalServerError,
		}
	}
	result := usecases.UserDetailResponse{
		ID:       ID,
		Username: users.Name,
		Email:    users.Email,
		Phone:    users.Phone,
	}
	return result, nil
}
