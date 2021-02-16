package main

import (
	"flag"
	"github.com/teratron/seabattle/cmd/seabattle/handler"
	"github.com/teratron/seabattle/pkg/server"
	"log"
	"os"

	"github.com/teratron/seabattle/pkg/config"
)

func init() {
	if err := os.Setenv("PORT", "8080"); err != nil {
		log.Println(err)
	}
}

func main() {
	cfg := config.New()
	flag.StringVar(&cfg.Addr, "addr", "localhost:8081", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./web/static", "Path to static assets")
	flag.Parse()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Инициализируем новую структуру с зависимостями приложения.
	/*app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}*/

	srv := server.New(cfg.Addr)

	srv.HandleFunc("/", handler.Home)
	srv.HandleFunc("/about", handler.About)
	srv.HandleFunc("/error", handler.Error)
	srv.HandleFileServer("./web/static")

	log.Fatal(srv.Run())
}
