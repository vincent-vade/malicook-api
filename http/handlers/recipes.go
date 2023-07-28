package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"recipes-api/domain/core"
	"recipes-api/domain/recipes"
	"recipes-api/infra/sqlc"
)

type RecipeHandler struct {
	recipeService recipes.IRecipeService
}

func NewRecipeHandler(rs recipes.IRecipeService) *RecipeHandler {
	return &RecipeHandler{
		recipeService: rs,
	}
}

func (rh *RecipeHandler) CreateRecipeHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var r sqlc.Recipe
		err := json.NewDecoder(request.Body).Decode(&r)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}

		er := rh.recipeService.NewRecipe(ctx, r)
		if er != nil {
			return
		}
		ctx.Render.JSON(writer, 200, map[string]interface{}{"msg": "ok"})
	}
}

func (rh *RecipeHandler) FindRecipeHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, _ := uuid.Parse(chi.URLParam(request, "id"))
		u := rh.recipeService.FindRecipe(ctx, id)

		err := ctx.Render.JSON(writer, 200, u)
		if err != nil {
			return
		}
	}

}

func (rh *RecipeHandler) ListRecipesHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		listRecipes, err := rh.recipeService.ListRecipes(ctx)
		if err != nil {
			http.Error(writer, err.Error(), 404)
		}
		ctx.Render.JSON(writer, 200, listRecipes)
	}
}
