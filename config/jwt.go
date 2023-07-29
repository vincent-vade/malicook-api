package config

import (
	"github.com/go-chi/jwtauth/v5"
	"os"
)

var TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
