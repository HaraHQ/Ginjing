// Controllers/Authentication/list.go

package Authentication

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtBody struct {
	Username string `json:"username"`
}

func Login(username string, password string) (*string, error) {
	if username == "admin" && password == "admin" {
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		signedToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			return nil, err
		}

		return &signedToken, nil
	}
	return nil, errors.New("username not found")
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return jwtToken, nil
}
