package tokenutil

import (
	"fmt"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clames)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return t, err
}

func CreateRefreshToken(user model.User, secret string, expire int) (string, error) {
	claimsRefresh := &domain.JwtCustomClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expire)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	rt, err := token.SignedString([]byte(secret))

	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIdFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("invalid Token")
	}

	return claims["id"].(string), nil
}
