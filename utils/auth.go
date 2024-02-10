package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "shreyash"

func GenerateJWTAuthToken(userId int64, email string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    userId,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func keyFuncHandler(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("wrong signing method")
	}
	return []byte(secret), nil
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, keyFuncHandler)
	if err != nil {
		return -1, errors.New("invalid Token: " + err.Error())
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, errors.New("invalid claims")
	}

	return int64(claims["id"].(float64)), nil
}
