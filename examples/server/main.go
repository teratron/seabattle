package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/teratron/seabattle/pkg/router"
)

func main() {
	// New создаём новое приложение
	srv := router.New()

	srv.HandleEntry()
	srv.HandleFunc("/test", test)
	srv.HandleFile("./web/static")

	// Run запускает приложение
	log.Fatal(srv.Start())
}

func test(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Test Page")
	//http.Redirect(w, r, "/", 302)
}
