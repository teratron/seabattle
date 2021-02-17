package main

import (
	"log"
	"os"

	"github.com/teratron/seabattle/cmd/seabattle/app"
)

func init() {
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Println(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	//
	a := app.New()

	//
	a.Server("localhost:8081")
}
