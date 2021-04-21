package db

import (
	"fmt"
	"log"

	"../models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const (
	user     = "postgres"
	password = "admin123"
	dbname   = "events"
	host     = "localhost"
	port     = "5432"
)

func GetConection() *gorm.DB {
	if db != nil {
		return db
	}
	conexion := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = gorm.Open(postgres.Open(conexion), &gorm.Config{})
	if err != nil {
		log.Println("Error en la conexion....")
		panic(err)
	}
	return db
}

func Migra() {
	_ = GetConection()
	db.AutoMigrate(&models.User{}, &models.Event{}, &models.Accesskey{})

}
