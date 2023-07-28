package categories

import (
	"context"
	"github.com/google/uuid"
	"recipes-api/infra/sqlc"
)

type CategoryRepository interface {
	Create(c context.Context, name string) error
	List(c context.Context) ([]sqlc.Category, error)
	GetByID(c context.Context, id uuid.UUID) (sqlc.Category, error)
	Edit(c context.Context, category sqlc.Category) error
	Remove(c context.Context, id uuid.UUID) error
}
