package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/goschool/crud/api"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", api.HandleHelloWorld)
	r.Post("/echo", api.HandleEchoUser)
	return r
}
