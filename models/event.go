package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"descripcion"`
	Keyword     string `json:"keyword"`
	Nivel       string `json:"nivel"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
