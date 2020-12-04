package database

import (
	"log"

	"github.com/ruyjfs/example-golang/config"
	"github.com/ruyjfs/example-golang/models"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.Db()
	db.AutoMigrate(
		&models.User{},
	)
	log.Printf("Migrate: Success")
}
