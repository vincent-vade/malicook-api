package handlers

import (
	"net/http"
	"recipes-api/domain/categories"
	"recipes-api/domain/core"
	"recipes-api/infra/sqlc"
)

func CreateCategoryHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		insertedUser := categories.NewCategory(ctx)
		ctx.Render.JSON(writer, 200, insertedUser)
	}

}

type UsersWithCount struct {
	categories []sqlc.Category
	count      int64
}

func ListCategoriesHandler(ctx *core.ContextHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		listCategories := categories.ListCategories(ctx)
		cnt := categories.CountUsers(ctx)
		ctx.Render.JSON(writer, 200, map[string]interface{}{"meta": cnt, "data": listCategories})
	}
}
