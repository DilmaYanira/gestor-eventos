package rutas

import "github.com/gorilla/mux"

var router *mux.Router

func RegistroRutas() *mux.Router {
	router = mux.NewRouter()
	routesUsers()
	routesEvents()
	routesAutenticar()
	routesAccessKey()
	return router
}
