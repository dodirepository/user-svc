package handlers

import (
	"net/http"

	render "github.com/dodirepository/user-svc/pkg"
)

type Healthcek struct{}

func (h Healthcek) Healthcek(w http.ResponseWriter, r *http.Request) {
	render.Render("ok", http.StatusOK, w)
}
