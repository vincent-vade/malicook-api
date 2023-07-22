package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"recipes-api/domain/core"
	"recipes-api/domain/users"
)

type UserHandler struct {
	userService users.IUserService
}

func NewUserHandler(us users.IUserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (uh *UserHandler) CreateUserHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := uh.userService.NewUser(ctx)
		if err != nil {
			return
		}
		ctx.Render.JSON(writer, 200, "de")
	}
}

func (uh *UserHandler) FindUserHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, _ := uuid.Parse(chi.URLParam(request, "id"))
		u := uh.userService.FindUser(ctx, id)

		err := ctx.Render.JSON(writer, 200, u)
		if err != nil {
			return
		}
	}

}

func (uh *UserHandler) ListUsersHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		listUsers, err := uh.userService.ListUsers(ctx)
		if err != nil {
			http.Error(writer, err.Error(), 404)
		}
		ctx.Render.JSON(writer, 200, listUsers)
	}
}
