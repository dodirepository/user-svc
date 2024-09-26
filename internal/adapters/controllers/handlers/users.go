package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	render "github.com/dodirepository/common-lib"
	domain "github.com/dodirepository/user-svc/internal/domain/usecases"
	"github.com/go-playground/validator/v10"
)

type UsersHandlersController struct {
	Usecase domain.UserUsecaseInterface
}

func (u UsersHandlersController) Create(w http.ResponseWriter, r *http.Request) {
	payload := domain.UserCreate{}
	err := render.ParseBody(r, &payload)
	if err != nil {
		render.Render(domain.ErrorResponse{
			Message: "Failed To Decode Payload",
		}, http.StatusUnprocessableEntity, w)
		return
	}

	validate := validator.New()
	trans := render.TranslatorValidatorIDN(validate)
	err = validate.Struct(payload)
	errs := render.TranslateError(err, trans)
	if errs != nil {
		render.Render(domain.ErrorResponse{
			Message: fmt.Sprintf("%v", errs),
		}, http.StatusBadRequest, w)
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
	err := render.ParseBody(r, &payload)
	if err != nil {
		render.Render(domain.ErrorResponse{
			Message: "Failed To Decode Payload",
		}, http.StatusUnprocessableEntity, w)
		return
	}
	validate := validator.New()
	trans := render.TranslatorValidatorIDN(validate)
	err = validate.Struct(payload)
	errs := render.TranslateError(err, trans)
	if errs != nil {
		render.Render(domain.ErrorResponse{
			Message: fmt.Sprintf("%v", errs),
		}, http.StatusBadRequest, w)
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
