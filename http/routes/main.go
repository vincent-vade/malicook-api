package routes

import (
	"github.com/go-chi/chi/v5"
	"recipes-api/domain/core"
)

func MakeRoutes(r *chi.Mux, cH *core.ContextHandler) {
	UsersRoute(r, cH)
	CategoriesRoute(r, cH)
	RecipesRoute(r, cH)
}
