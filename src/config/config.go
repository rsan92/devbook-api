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
	SecretKey          = make([]byte, 64)
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
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_NOME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	if len(SecretKey) == 0 {
		log.Fatal("Falha ao obter secret key")
	}
}
