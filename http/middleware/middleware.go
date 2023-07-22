package middleware

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MakeMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
}
