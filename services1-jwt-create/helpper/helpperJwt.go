package helpper

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte("uwhiuwhoqwWYUGw234323@$%^&")

func GetJwt() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	Claims := token.Claims.(jwt.MapClaims)

	Claims["authorized"] = true
	Claims["client"] = "asliddin"
	Claims["aud"] = "asliddin.jwtgo.io"
	Claims["iss"] = "jwt.io"
	Claims["exp"] = time.Now().Add(time.Minute * 2).UTC()

	tokenString, err := token.SignedString(MySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
