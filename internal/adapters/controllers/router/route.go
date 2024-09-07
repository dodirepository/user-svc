package router

import (
	"net/http"

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
	root.HandleFunc("/health", healthCheckControler.Healthcek).Methods(http.MethodGet)

	in := root.PathPrefix("/in").Subrouter()
	in.Use(middleware.JWTAuthMiddleware)
	root.HandleFunc("/login", userController.Login).Methods(http.MethodPost)
	users := in.PathPrefix("/users").Subrouter()

	users.HandleFunc("",
		userController.Create).Methods(http.MethodPost)
	users.HandleFunc("/profile",
		userController.Detail).Methods(http.MethodGet)
	return rtr.router
}
