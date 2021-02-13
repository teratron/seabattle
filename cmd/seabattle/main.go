package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/teratron/seabattle/cmd/seabattle/handler"
	"github.com/teratron/seabattle/pkg/router"
)

func init() {
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Println(err)
	}
}

type Config struct {
	Addr      string
	StaticDir string
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	// Используется функция http.NewServeMux() для инициализации нового рутера.
	r := router.NewRouter()

	r.HandleFunc("/", handler.Home)
	r.HandleFunc("/about", handler.About)
	r.HandleFunc("/error", handler.Error)
	r.HandleFileServer("./web/static")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
