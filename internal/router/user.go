package router

import (
	"net/http"
	"template/internal/handler"

	"github.com/go-chi/chi/v5"
)

func (rtr *router) UserRouter() http.Handler {
	userHandler := handler.NewUserHandler()
	router := chi.NewRouter()

	router.Post("/", userHandler.Create)
	router.Get("/", userHandler.List)
	router.Get("/{id}", userHandler.Get)
	router.Put("/{id}", userHandler.Update)

	return router
}
