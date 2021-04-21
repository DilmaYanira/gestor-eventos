package db

import "../models"

func CreateEvent(body models.Event) (event models.Event) {
	db := GetConection()
	db.Create(&body)
	db.Last(&event)
	return

}

func ListarEvents() (events []models.Event) {
	db := GetConection()
	db.Find(&events)
	return
}

func ListarEventsPublics() (events []models.Event) {
	db := GetConection()
	db.Find(&events, "nivel = ?", "publico")
	return
}

func GetEventById(id uint) (event models.Event) {
	db := GetConection()
	db.Find(&event, "id = ? ", id)
	return
}
