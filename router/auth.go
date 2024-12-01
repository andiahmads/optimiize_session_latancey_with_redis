package router

import (
	"gored/middleware"
	"gored/pkg/auth"

	"github.com/go-chi/chi/v5"
)

func UserRouters() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/login", auth.LoginHandler)
	r.With(middleware.Authorization).Get("/product", auth.GetProductHandler)

	return r
}
