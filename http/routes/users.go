package routes

import (
	"github.com/go-chi/chi/v5"
	"recipes-api/domain/core"
	"recipes-api/domain/users"
	"recipes-api/http/handlers"
	"recipes-api/infra/sqlc/repository"
)

func UsersRoute(r chi.Router, cH *core.ContextHandler) {
	ur := repository.NewUserRepository(cH.SqlcCtx)
	us := users.NewUserService(ur)
	uh := handlers.NewUserHandler(us)

	r.Route("/users", func(router chi.Router) {
		router.Post("/", uh.CreateUserHandler(cH))
		router.Get("/", uh.ListUsersHandler(cH))
		router.Get("/{id}", uh.FindUserHandler(cH))
	})
}
