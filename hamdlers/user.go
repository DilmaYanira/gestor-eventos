package hamdlers

import (
	"net/http"

	"../io/request"
	"../io/response"
	"../models"
	"../services"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.User
	request.Json(r, &body)
	msg, user, err := services.CreateUser(body)
	if err != nil {
		response.Error(msg, http.StatusBadRequest, w)
		return
	}
	response.Json(user, http.StatusOK, w)

}

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	user := services.GetUserByEmail(email)
	response.Json(user, http.StatusOK, w)
}
