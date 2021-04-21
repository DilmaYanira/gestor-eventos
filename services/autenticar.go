package services

import (
	"log"

	"../db"
	"../jwt"
	"../models"
	"golang.org/x/crypto/bcrypt"
)

func Login(body models.User) (loginResp models.Login, mensaje string, err error) {

	if len(body.Email) <= 0 || len(body.Password) <= 0 {
		mensaje = "el Email y el password son obligatorios"
		return
	}

	user := db.GetUserByEmail(body.Email)
	log.Println(user)
	userPass := []byte(user.Password)
	bodyPass := []byte(body.Password)

	err = bcrypt.CompareHashAndPassword(userPass, bodyPass)
	log.Println(err)
	if err != nil {
		mensaje = "Error en el password"
		return
	}
	token, err := jwt.GenerarToken(user)
	if err != nil {
		mensaje = "Error generando token"
		return
	}
	user.Password = ""
	loginResp = models.Login{Token: token, User: user}
	return

}
