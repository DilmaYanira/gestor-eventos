package hamdlers

import (
	"net/http"

	"../io/response"
	"../services"
)

func CreateAccessKey(w http.ResponseWriter, r *http.Request) {
	respkey := services.CreateAccessKey()
	response.Json(respkey, http.StatusOK, w)
}
