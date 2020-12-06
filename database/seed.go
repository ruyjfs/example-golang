package database

import (
	"log"

	"github.com/ruyjfs/example-golang/database/seeds"
)

func Seeder() {
	log.Printf("Seed: Start")
	seeds.Perfils()
	log.Printf("Seed: Success")
}
