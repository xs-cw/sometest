package handler

import (
	"demo-go/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func NewToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	return token.SignedString([]byte(config.GetSecretKey()))
}
