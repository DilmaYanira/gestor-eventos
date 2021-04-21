package services

import (
	"errors"
	"regexp"

	"../crypto"
	"../db"
	"../models"
)

func CreateUser(body models.User) (mensaje string, user models.User, err error) {
	err = validateUser(body)
	if err != nil {
		mensaje = err.Error()
		return
	}
	existingUser := db.GetUserByEmail(body.Email)
	if existingUser.ID > 0 {
		mensaje = "Email Duplicado"
		return
	}
	pass, err := crypto.Ecrypting(body.Password)
	if err != nil {
		mensaje = "Error de encriptacion del pasword"
		return
	}
	body.Password = pass
	user = db.CreateUser(body)
	return
}

func validateUser(user models.User) (err error) {
	if len(user.Name) < 5 {
		err = errors.New("El nombre debe superar los 5 caracteres")
		return
	}
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(user.Email) < 6 || len(user.Email) > 255 || !regex.MatchString(user.Email) {
		err = errors.New("Error en el formato del email")
		return
	}
	if len(user.Password) < 6 {
		err = errors.New("El pasword debe tener 6 o mas caracteres ")
		return
	}
	return
}

func GetUserByEmail(email string) (user models.User) {
	user = db.GetUserByEmail(email)
	return
}
