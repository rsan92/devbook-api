package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.Carregar()
	fmt.Printf("Executando API na porta [%v]...\n", config.PortaAPI)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PortaAPI), r))
}
