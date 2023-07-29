package routes

import (
	"github.com/go-chi/chi/v5"
	"recipes-api/domain/core"
	"recipes-api/domain/recipes"
	"recipes-api/http/handlers"
	"recipes-api/infra/sqlc/repository"
)

func RecipesRoute(r chi.Router, cH *core.ContextHandler) {
	rr := repository.NewRecipeRepository(cH.SqlcCtx)
	rs := recipes.NewRecipeService(rr)
	rh := handlers.NewRecipeHandler(rs)

	r.Route("/recipes", func(router chi.Router) {
		router.Post("/", rh.CreateRecipeHandler(cH))
		router.Get("/", rh.ListRecipesHandler(cH))
		router.Get("/{id}", rh.FindRecipeHandler(cH))
	})
}
