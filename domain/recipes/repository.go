package recipes

import (
	"context"
	"github.com/google/uuid"
	"recipes-api/infra/sqlc"
)

type RecipeRepository interface {
	Create(c context.Context, recipe sqlc.Recipe) (sqlc.Recipe, error)
	List(c context.Context) ([]sqlc.Recipe, error)
	GetByID(c context.Context, id uuid.UUID) (sqlc.Recipe, error)
	Edit(c context.Context, recipe sqlc.Recipe) error
	Remove(c context.Context, id uuid.UUID) error
}
