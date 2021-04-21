package rutas

import (
	"../hamdlers"
	"../midlewars"
)

func routesAccessKey() {
	router.HandleFunc("/Key", midlewars.Admin(midlewars.Auth(hamdlers.CreateAccessKey))).Methods("GET")
}
