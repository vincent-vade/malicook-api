package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"recipes-api/config"
	"recipes-api/domain/core"
)

func MakePrivateRoutes(r *chi.Mux, cH *core.ContextHandler) {
	r.Group(func(router chi.Router) {
		router.Use(jwtauth.Verifier(config.TokenAuth))
		router.Use(jwtauth.Authenticator)
		UsersRoute(router, cH)
		CategoriesRoute(router, cH)
		RecipesRoute(router, cH)
	})
}
