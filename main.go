package main

import (
	database "github.com/matiaslamela/go-boilerplate/src/database"
	enviroment "github.com/matiaslamela/go-boilerplate/src/env"
	router "github.com/matiaslamela/go-boilerplate/src/server"
)

func main() {
	enviroment.GetEnvVars()
	database.Connect()
	app := router.Server{}
	app.Start()
}
