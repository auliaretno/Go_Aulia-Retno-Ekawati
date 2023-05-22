package main

import (
	"code_competence/config"
	"code_competence/routes"
)

func main() {
	// start the server, and log if it fails
	config.InitDB()
	config.InitialMigration()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
