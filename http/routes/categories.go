package routes

import (
	"github.com/go-chi/chi/v5"
	"recipes-api/domain/core"
	"recipes-api/http/handlers"
)

func CategoriesRoute(r chi.Router, cH *core.ContextHandler) {
	r.Route("/categories", func(router chi.Router) {
		router.Post("/", handlers.CreateCategoryHandler(cH))
		router.Get("/", handlers.ListCategoriesHandler(cH))
	})
}
