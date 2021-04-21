package db

import (
	"../models"
)

func CreateAccessKey(body models.Accesskey) (access models.Accesskey) {
	db := GetConection()
	db.Create(&body)
	db.Last(&access)
	return

}
func ConsultarAccessKey(key string) (access models.Accesskey) {
	db := GetConection()
	db.Find(&access, "key = ?", key)
	return

}
