package tokenutil

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/krishnatrea/cdp/database/model"
	"github.com/krishnatrea/cdp/domain"
)

func CreateAccessToken(user model.User, secret string, expire int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	clames := &domain.JwtCustomClaims{
		Name: user.Firstname,
		ID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, clames)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, err
}
