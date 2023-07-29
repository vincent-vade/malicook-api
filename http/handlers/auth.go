package handlers

import (
	"encoding/json"
	"net/http"
	"recipes-api/domain/auth"
	"recipes-api/domain/core"
	"recipes-api/infra/sqlc"
)

type AuthHandler struct {
	authService auth.IAuthService
}

func NewAuthHandler(as auth.IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: as,
	}
}

func (ah *AuthHandler) LoginHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var loginBody auth.LoginBody
		if request.Body == nil {
			http.Error(writer, "Please send a request body", http.StatusBadRequest)
		}

		err := json.NewDecoder(request.Body).Decode(&loginBody)
		if err != nil {
			http.Error(writer, err.Error(), 500)
		}

		tkn := ah.authService.Login(ctx, loginBody)
		// call ah.authService.login(email, password)
		// render json with 200 success
		ctx.Render.JSON(writer, 200, map[string]interface{}{"token": tkn, "success": true})
	}
}

func (ah *AuthHandler) RegisterHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var registerBody sqlc.User
		if request.Body == nil {
			http.Error(writer, "Please send a request body", http.StatusBadRequest)
		}

		err := json.NewDecoder(request.Body).Decode(&registerBody)
		if err != nil {
			http.Error(writer, err.Error(), 500)
		}

		ah.authService.Register(ctx, registerBody)

		ctx.Render.JSON(writer, 200, map[string]interface{}{"msg": "ok"})
	}
}

func (ah *AuthHandler) LogoutHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx.Render.JSON(writer, 200, map[string]interface{}{"msg": "ok"})
	}
}

func (ah *AuthHandler) VerifyTokenHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx.Render.JSON(writer, 200, map[string]interface{}{"msg": "ok"})
	}
}
