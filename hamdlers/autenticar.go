package hamdlers

import (
	"net/http"

	"../io/request"
	"../io/response"
	"../models"
	"../services"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	err := request.Json(r, &body)
	if err != nil {
		response.Error("Error al decodificar el body", http.StatusBadRequest, w)
		return
	}
	loginResp, msg, err := services.Login(body)
	if msg != "" || err != nil {
		response.Error(msg, http.StatusBadRequest, w)
		return
	}
	response.Json(loginResp, http.StatusOK, w)

}
