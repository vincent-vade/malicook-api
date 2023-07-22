package categories

import (
	"log"
	"recipes-api/domain/core"
	"recipes-api/infra/sqlc"
)

func NewCategory(ctx *core.ContextHandler) sqlc.Category {
	insertedCategory, err := ctx.SqlcCtx.CreateCategory(ctx.Ctx, "dede")
	if err != nil {
		log.Fatal(err)
	}
	return insertedCategory
}

func ListCategories(ctx *core.ContextHandler) []sqlc.Category {
	listCategories, err := ctx.SqlcCtx.ListCategories(ctx.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	return listCategories
}

func CountUsers(ctx *core.ContextHandler) int64 {
	count, err := ctx.SqlcCtx.CountCategory(ctx.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
