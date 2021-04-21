package services

import (
	"errors"

	"../db"
	"../models"
)

func CreateEvent(body models.Event) (mensaje string, event models.Event, err error) {
	if len(event.Title) == 0 {
		err = errors.New("Tiltulo es obligatorio")
		return
	}
	if len(event.Description) == 0 {
		err = errors.New("Descripcion es obligatoria")
		return
	}
	if len(event.Nivel) == 0 {
		err = errors.New("Nivel es un campo obligatorio")
	}

	event = db.CreateEvent(body)
	return

}
func ListarEvents(acessKey []string) (events []models.Event, err error) {
	if len(acessKey) > 0 {
		if ValidarAccessKey(acessKey[0]) {
			events = db.ListarEvents()
		} else {
			err = errors.New("Access key no autorizada")
		}

	} else {
		events = db.ListarEventsPublics()
	}
	return
}

func GetEventById(id uint, acessKey []string) (event models.Event, err error) {
	if len(acessKey) > 0 {
		if !ValidarAccessKey(acessKey[0]) {
			err = errors.New("Access key no autorizada")
			return
		}
		event = db.GetEventById(id)
		return
	}
	event = db.GetEventById(id)
	if event.Nivel != "publico" {
		err = errors.New("No está autorizado para ver este evento")
		return
	}

	return
}
