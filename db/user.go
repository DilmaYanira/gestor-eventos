package db

import "../models"

func CreateUser(body models.User) (user models.User) {
	db := GetConection()
	db.Create(&body)
	db.Last(&user)
	return
}
func GetUserByEmail(email string) (user models.User) {
	db := GetConection()
	db.Find(&user, "email = ?", email)
	return
}
