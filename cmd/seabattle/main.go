package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/teratron/seabattle/cmd/seabattle/handler"
)

func init() {
	err := os.Setenv("PORT", "8080")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// Используется функция http.NewServeMux() для инициализации нового рутера.
	mux := http.NewServeMux()

	// HomeHandler регистрируется как обработчик для URL-шаблона "/".
	mux.HandleFunc("/", handler.Home)

	// AboutHandler регистрируется как обработчик для URL-шаблона "/about".
	mux.HandleFunc("/about", handler.About)

	mux.HandleFunc("/error", handler.Error)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
