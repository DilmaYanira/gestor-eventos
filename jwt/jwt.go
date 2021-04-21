package jwt

import (
	"errors"
	"strings"
	"time"

	"../models"
	jwt "github.com/dgrijalva/jwt-go"
)

var JWT_KEY = []byte("dilma")
var UserEmail string
var UserID string

func GenerarToken(user models.User) (string, error) {
	payload := jwt.MapClaims{
		"email": user.Email,
		"name":  user.Name,
		"id":    user.ID,
		"rol":   user.Rol,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(JWT_KEY)
	return token, err
}

func ProcesarToken(authHeader string) (*models.Claim, error) {
	claims := &models.Claim{}
	if len(authHeader) <= 0 {
		return claims, errors.New("No auth token present")
	}
	splitToken := strings.Split(authHeader, "Bearer")
	if len(splitToken) != 2 {
		return claims, errors.New("Invalid token format")
	}
	tokenStr := strings.TrimSpace(splitToken[1])
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil {
		return claims, err
	}

	return claims, nil
}
