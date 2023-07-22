package repository

import (
	"context"
	"github.com/google/uuid"
	"log"
	"recipes-api/domain/users"
	sqlc2 "recipes-api/infra/sqlc"
)

type userRepository struct {
	sqlc *sqlc2.Queries
}

func NewUserRepository(sqlc *sqlc2.Queries) users.UserRepository {
	return &userRepository{
		sqlc: sqlc,
	}
}

func (ut *userRepository) List(c context.Context) ([]sqlc2.User, error) {
	listUsers, err := ut.sqlc.ListUsers(c)
	return listUsers, err
}

func (ut *userRepository) Create(c context.Context, user *sqlc2.User) error {
	_, err := ut.sqlc.CreateUser(c, sqlc2.CreateUserParams{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})

	return err
}

func (ut *userRepository) GetByID(c context.Context, id uuid.UUID) (sqlc2.User, error) {
	usr, err := ut.sqlc.GetUser(c, id)
	if err != nil {
		log.Fatal(err)
	}
	return usr, err
}
