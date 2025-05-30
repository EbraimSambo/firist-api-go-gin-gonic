package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ebraimsambo"
	password = "Deezycheezy@2"
	dbname   = "go"
)

func ConnectDB() (*sql.DB, error) {
	psqInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Banco de dados conectado: %s\n", dbname)
	return db, nil
}
