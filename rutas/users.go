package rutas

import (
	"../hamdlers"
)

func routesUsers() {
	router.HandleFunc("/users", hamdlers.CreateUser).Methods("POST")
	router.HandleFunc("/byemail/{email}", hamdlers.GetUserByEmail).Methods("GET")

}
