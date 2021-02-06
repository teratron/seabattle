package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/teratron/seabattle/pkg/router"
)

func init() {
	err := os.Setenv("PORT", "8080")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	fmt.Println("Sea Battle")

	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
