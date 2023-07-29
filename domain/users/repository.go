package users

import (
	"context"
	"github.com/google/uuid"
	"recipes-api/infra/sqlc"
)

type UserRepository interface {
	Create(c context.Context, user sqlc.User) (uuid.UUID, error)
	List(c context.Context) ([]sqlc.User, error)
	GetByID(c context.Context, id uuid.UUID) (sqlc.User, error)
	GetByEmail(c context.Context, email string) (sqlc.User, error)
}
