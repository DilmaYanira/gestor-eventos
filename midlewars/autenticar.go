package midlewars

import (
	"errors"
	"net/http"

	"../jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		_, err := jwt.ProcesarToken(header)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	}
}

func Admin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		claims, _ := jwt.ProcesarToken(header)
		if claims.Rol != "Administrador" {
			w.WriteHeader(http.StatusUnauthorized)
			err := errors.New("El usuario no es administrador")
			w.Write([]byte(err.Error()))
			return
		}
		next.ServeHTTP(w, r)
	}
}
