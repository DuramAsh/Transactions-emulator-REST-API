package utility

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	emulator "github.com/duramash/constanta-emulator-task/src/models"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

const (
	signingKey = "asdjfluo13270894hasfdh9210*&()asdfasd7"
)

func GenerateToken(username, password string) (string, error) {
	expectedPassword, ok := emulator.Users[username]
	if !ok || expectedPassword != password {
		return "", fmt.Errorf("wrong credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		username,
	})
	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", fmt.Errorf("wrong type of token claims")
	}
	return claims.Username, nil
}

func IdentifyUser(c *gin.Context) error {
	header := c.GetHeader("Authorization")

	if len(header) == 0 {
		return fmt.Errorf("no authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return fmt.Errorf("wrong header")
	}

	username, err := ParseToken(headerParts[1])
	if err != nil {
		return err
	}

	c.Set("username", username)
	return nil
}
