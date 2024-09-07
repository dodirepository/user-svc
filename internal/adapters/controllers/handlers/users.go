package handlers

import (
	"net/http"
	"strconv"

	domain "github.com/dodirepository/user-svc/internal/domain/usecases"
	parser "github.com/dodirepository/user-svc/pkg"
	render "github.com/dodirepository/user-svc/pkg"
)

type UsersHandlersController struct {
	Usecase domain.UserUsecaseInterface
}

func (u UsersHandlersController) Create(w http.ResponseWriter, r *http.Request) {
	payload := domain.UserCreate{}
	err := parser.ParseBody(r, &payload)
	if err != nil {
		render.Render(domain.ErrorResponse{
			Message: "Failed To Decode Payload",
		}, http.StatusUnprocessableEntity, w)
		return
	}

	resperr := u.Usecase.Insert(payload)
	if resperr != nil {
		render.Render(resperr, resperr.Status, w)
		return
	}

	render.Render("ok", http.StatusOK, w)
}

func (u UsersHandlersController) Login(w http.ResponseWriter, r *http.Request) {
	payload := domain.UserLogin{}
	err := parser.ParseBody(r, &payload)
	if err != nil {
		render.Render(domain.ErrorResponse{
			Message: "Failed To Decode Payload",
		}, http.StatusUnprocessableEntity, w)
		return
	}

	result, resperr := u.Usecase.Login(payload)
	if resperr != nil {
		render.Render(resperr, resperr.Status, w)
		return
	}

	render.Render(result, http.StatusOK, w)
}
func (u UsersHandlersController) Detail(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-User-ID")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		render.Render(domain.ErrorResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError, w)
		return
	}

	result, resperr := u.Usecase.Detail(userID)
	if resperr != nil {
		render.Render(resperr, resperr.Status, w)
		return
	}

	render.Render(result, http.StatusOK, w)
}
