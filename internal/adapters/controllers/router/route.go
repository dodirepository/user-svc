package router

import (
	"net/http"

	lib "github.com/dodirepository/common-lib"
	"github.com/dodirepository/user-svc/internal/middleware"
	"github.com/gorilla/mux"
)

type router struct {
	router *mux.Router
}

// NewRouter :nodoc:
func NewRouter() Router {
	return &router{
		router: mux.NewRouter(),
	}
}

// Route :nodoc:
func (rtr *router) Route() *mux.Router {
	root := rtr.router.PathPrefix("/").Subrouter()
	internal := root.PathPrefix("/in").Subrouter()
	v1 := internal.PathPrefix("/v1").Subrouter()
	v1.Use(middleware.JWTAuthMiddleware)

	//internal group
	internal.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	internal.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		lib.Render("ok", http.StatusOK, w)
	}).Methods(http.MethodGet)
	internal.Handle("/validate", middleware.JWTAuthMiddleware(http.HandlerFunc(userController.Detail))).Methods(http.MethodGet)

	//v1/users group
	users := v1.PathPrefix("/users").Subrouter()

	internal.HandleFunc("/register",
		userController.Create).Methods(http.MethodPost)
	users.HandleFunc("/profile",
		userController.Detail).Methods(http.MethodGet)
	return rtr.router
}
