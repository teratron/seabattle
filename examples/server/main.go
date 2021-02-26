package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teratron/seabattle/pkg/router"
)

func main() {
	// New создаём новое приложение
	r := router.New()

	r.HandleEntry()
	r.HandleFunc("/test", test)
	r.HandleFile("./web/static")

	// Run запускает приложение
	log.Fatal(r.Run())
}

func test(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Test Page")
}
