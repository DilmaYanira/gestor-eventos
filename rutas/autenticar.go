package rutas

import (
	"../hamdlers"
)

func routesAutenticar() {
	router.HandleFunc("/Login", hamdlers.Login).Methods("POST")
}
