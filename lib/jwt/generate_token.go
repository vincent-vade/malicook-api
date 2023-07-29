package jwt

import (
	"github.com/lestrrat-go/jwx/v2/jwt"
	"recipes-api/config"
)

func GenerateToken(payload map[string]interface{}) string {
	_, tokenString, _ := config.TokenAuth.Encode(payload)
	return tokenString
}

func DecodeToken(token string) jwt.Token {
	tokenString, _ := config.TokenAuth.Decode(token)
	return tokenString
}
