package main

import (
	"github.com/ruyjfs/example-golang/database"
	"github.com/ruyjfs/example-golang/routes"
)

func main() {
	database.Migrate()
	routes.Run()
}
