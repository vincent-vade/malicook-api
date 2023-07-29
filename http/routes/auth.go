package routes

import (
	"github.com/go-chi/chi/v5"
	"recipes-api/domain/auth"
	"recipes-api/domain/core"
	"recipes-api/http/handlers"
	"recipes-api/infra/sqlc/repository"
)

func AuthRoute(r *chi.Mux, cH *core.ContextHandler) {
	ur := repository.NewUserRepository(cH.SqlcCtx)
	us := auth.NewAuthService(ur)
	ah := handlers.NewAuthHandler(us)

	r.Route("/auth", func(router chi.Router) {
		router.Post("/login", ah.LoginHandler(cH))
		router.Post("/register", ah.RegisterHandler(cH))
		router.Post("/logout", ah.LogoutHandler(cH))
		router.Get("/verify/{token}", ah.VerifyTokenHandler(cH))
	})
}
