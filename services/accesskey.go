package services

import (
	"math/rand"
	"strings"

	"../db"

	"../models"
)

func CreateAccessKey() (access models.Accesskey) {
	var body models.Accesskey
	body.Key = generarKey()
	body.Status = "Activo"
	access = db.CreateAccessKey(body)
	return

}

func ValidarAccessKey(key string) (resp bool) {

	access := db.ConsultarAccessKey(key)

	if access.Key != "" {
		resp = true
		return
	}
	return
}

func generarKey() string {
	caracteres := "0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z"
	arrayCaracteres := strings.Split(caracteres, ",")
	var key string
	for i := 0; i < 8; i++ {

		indice := rand.Intn(36)
		key = key + arrayCaracteres[indice]

	}
	return key

}
