package banco

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"api/src/config"
)

func Conectar() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoBanco)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
