package middlewares

import (
	"log"
	"net/http"

	"api/src/autenticacao"
	"api/src/respostas"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{
			log.Printf("\n %s - %s - %s", r.Method, r.RequestURI, r.Host)
			next(w, r)
		}
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			respostas.Erro(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
