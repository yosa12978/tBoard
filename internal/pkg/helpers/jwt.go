package helpers

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/yosa12978/tBoard/internal/pkg/models"
)

func ParseJWT(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

func NewJWT(account models.Account) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": account.Username,
			"id":       account.ID,
			"role":     account.Role,
		})
	return token.SignedString(key)
}
