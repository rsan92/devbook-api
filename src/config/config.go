package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	PortaAPI           = 0
)

func Carregar() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	p, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		p = 9000
	}

	PortaAPI = p
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf-8&parseTime=True&loc=Local", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))
}
