package auth

import (
	"recipes-api/infra/sqlc"
)

type LoginBody struct {
	Email    string
	Password string
}
type LoginAuth struct {
	token string
}

type RegisterBody struct {
	sqlc.User
}
