package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/teratron/seabattle/cmd/seabattle/handler"
	"github.com/teratron/seabattle/pkg/router"
)

//
type ServeMux struct {
	http.ServeMux
}

func init() {
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Println(err)
	}
}

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера.
	mux := router.NewServeMux()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/about", handler.About)
	mux.HandleFunc("/error", handler.Error)
	mux.HandleFileServer("./web/static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
