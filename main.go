package main

import (
	"log"
	"net/http"

	"./db"
	"./rutas"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Addr   string
}

func (s *Server) Initialize(addr string) {
	s.Router = rutas.RegistroRutas()
	s.Addr = addr
}

func (s *Server) Run() {
	log.Print("Ejecutando por el puerto", s.Addr)
	http.Handle("/", s.Router)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}
func main() {
	db.Migra()
	server := Server{}
	server.Initialize(":5000")
	server.Run()
}
