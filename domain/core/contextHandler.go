package core

import (
	"context"
	"github.com/unrolled/render"
	"recipes-api/infra/sqlc"
)

type ContextHandler struct {
	Ctx     context.Context
	SqlcCtx *sqlc.Queries
	Render  *render.Render
}

func NewContextHandler(ctx context.Context, sqlcCtx *sqlc.Queries, r *render.Render) *ContextHandler {
	return &ContextHandler{
		Ctx:     ctx,
		SqlcCtx: sqlcCtx,
		Render:  r,
	}
}
