package main

import (
	"log"

	"github.com/teratron/seabattle/pkg/router"
)

func main() {
	// New создаём новое приложение
	r := router.New()

	r.HandleEntry()
	r.HandleFile("./web/static")

	// Run запускает приложение
	log.Fatal(r.Run())
}
