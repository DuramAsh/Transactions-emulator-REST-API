package emulator

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var SigningKey = []byte("asdlfhalhd78103849hal;dl9i[8)S&lafdsjkfs")
