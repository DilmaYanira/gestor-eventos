package hamdlers

import (
	"net/http"
	"strconv"

	"../io/request"
	"../io/response"
	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var body models.Event
	request.Json(r, &body)
	_, even, _ := services.CreateEvent(body)
	response.Json(even, http.StatusOK, w)
}

func ListarEvents(w http.ResponseWriter, r *http.Request) {
	acess_key := r.URL.Query()["key"]
	events, err := services.ListarEvents(acess_key)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(events, http.StatusOK, w)
}
func GetEventById(w http.ResponseWriter, r *http.Request) {
	acess_key := r.URL.Query()["key"]
	vars := mux.Vars(r)
	id, errv := strconv.Atoi(vars["id"])
	if errv != nil {
		response.Error("Error en conversion de id", http.StatusBadRequest, w)
		return
	}
	event, err := services.GetEventById(uint(id), acess_key)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	response.Json(event, http.StatusOK, w)

}
