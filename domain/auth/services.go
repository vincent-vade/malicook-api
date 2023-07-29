package auth

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"recipes-api/domain/core"
	"recipes-api/domain/users"
	"recipes-api/infra/sqlc"
	"recipes-api/lib/jwt"
	"recipes-api/lib/mailer"
)

type AuthService struct {
	ur users.UserRepository
}

type IAuthService interface {
	Login(ch *core.ContextHandler, body LoginBody) string
	Register(ch *core.ContextHandler, body sqlc.User) uuid.UUID
	VerifyToken(ch *core.ContextHandler, token string)
	Logout()
}

func NewAuthService(ur users.UserRepository) IAuthService {
	return &AuthService{
		ur: ur,
	}
}

func (as AuthService) VerifyToken(ch *core.ContextHandler, token string) {
	//TODO implement me
	panic("implement me")
}

func (as AuthService) Login(c *core.ContextHandler, body LoginBody) string {
	// validate payload
	u, err := as.ur.GetByEmail(c.Ctx, body.Email)
	fmt.Printf("email %s", u.Email)
	if err != nil {
		log.Fatal("User not exist")
	}

	tokenAuth := jwt.GenerateToken(map[string]interface{}{"id": u.ID})

	return tokenAuth
}

func (as AuthService) Register(c *core.ContextHandler, body sqlc.User) uuid.UUID {
	// validate payload
	uId, err := as.ur.Create(c.Ctx, body)
	fmt.Printf("user %s \n", uId)
	if err != nil {
		log.Fatal("error when creating user")
	}

	usr, _ := as.ur.GetByID(c.Ctx, uId)
	tkn := jwt.GenerateToken(map[string]interface{}{"id": usr.ID})

	// update user with verifyToken
	// create endpoint for confirmation token
	// when confirm
	// update user

	mailer.SendMail(usr.Email, "Confirm account creation !", "confirmation.html", mailer.ConfirmationBody{Token: tkn, FullName: usr.FirstName + " " + usr.LastName})

	return uId
}

func (as AuthService) Logout() {
	//TODO implement me
	panic("implement me")
}
