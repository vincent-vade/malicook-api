package users

import (
	"github.com/google/uuid"
	"recipes-api/domain/core"
	"recipes-api/infra/sqlc"
	"recipes-api/lib/bcrypt"
)

type UserService struct {
	userRepository UserRepository
}

type IUserService interface {
	NewUser(ch *core.ContextHandler, u *sqlc.User) (uuid.UUID, error)
	FindUser(ch *core.ContextHandler, id uuid.UUID) sqlc.User
	FindUserByEmail(ch *core.ContextHandler, email string) sqlc.User
	ListUsers(ch *core.ContextHandler) ([]sqlc.User, error)
}

func NewUserService(ur UserRepository) IUserService {
	return &UserService{userRepository: ur}
}

func genPass(pass string, c chan string) string {
	hPass, _ := bcrypt.HashPassword(pass)
	c <- hPass
	return hPass
}

func chanPass(pass string) string {
	c := make(chan string)
	defer close(c)

	go genPass(pass, c)

	receivePass := <-c
	return receivePass
}

func (us UserService) NewUser(ctx *core.ContextHandler, u *sqlc.User) (uuid.UUID, error) {
	pass := chanPass(u.Password)
	cUsr := sqlc.User{
		Email:     u.Email,
		LastName:  u.LastName,
		FirstName: u.FirstName,
		Password:  pass,
	}
	x, err := us.userRepository.Create(ctx.Ctx, cUsr)

	return x, err
}

func (us UserService) FindUser(ch *core.ContextHandler, id uuid.UUID) sqlc.User {
	user, _ := us.userRepository.GetByID(ch.Ctx, id)
	return user
}

func (us UserService) FindUserByEmail(ch *core.ContextHandler, email string) sqlc.User {
	user, _ := us.userRepository.GetByEmail(ch.Ctx, email)
	return user
}

func (us UserService) ListUsers(ctx *core.ContextHandler) ([]sqlc.User, error) {
	listUsers, err := us.userRepository.List(ctx.Ctx)
	return listUsers, err
}
