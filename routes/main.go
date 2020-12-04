package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/ruyjfs/example-golang/controllers"
)

const defaultPort = "8080"

func Run() {
	index := new(controllers.Index)
	http.HandleFunc("/", index.Index)

	GraphQL()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
