package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/goschool/crud/api"
)

func SetupRoutes(userHandler api.UserHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", api.HandleHelloWorld)
	r.Post("/echo", api.HandleEchoUser)
	r.Post("/register", userHandler.HandlerRegisterUser)
	return r
}
