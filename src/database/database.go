package database

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	enviroment "github.com/matiaslamela/go-boilerplate/src/env"
)

type Connection struct {
	DB *sql.DB
}

var Conn Connection

func Connect() {
	db, err := sql.Open("postgres", enviroment.DB_URI)

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	Conn = Connection{
		DB: db,
	}
}
