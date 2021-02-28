package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teratron/seabattle/pkg/server"
)

func main() {
	// New создаём новое приложение
	srv := server.New()

	srv.HandleEntry()
	srv.HandleFunc("/test", test)
	srv.HandleFile("./web/static")

	// Run запускает приложение
	log.Fatal(srv.Start())
}

func test(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Test Page")
}
