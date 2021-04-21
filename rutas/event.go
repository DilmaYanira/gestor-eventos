package rutas

import (
	"../midlewars"

	"../hamdlers"
)

func routesEvents() {
	router.HandleFunc("/Events", midlewars.Admin(midlewars.Auth(hamdlers.CreateEvent))).Methods("POST")
	router.HandleFunc("/Events", hamdlers.ListarEvents).Methods("GET")
	router.HandleFunc("/Event/{id}", hamdlers.GetEventById).Methods("GET")

}
