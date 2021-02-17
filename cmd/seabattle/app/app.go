package app

import (
	"flag"
	"log"
	"sync"

	"github.com/teratron/seabattle/cmd/seabattle/handler"
	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
	"github.com/teratron/seabattle/pkg/server"
)

type Application struct {
	srv      *server.Server
	cfg      *config.Config
	log      *logger.Logger
	settings *settings
	mu       sync.Mutex
}

// New
func New() *Application {
	app := &Application{
		srv:      &server.Server{},
		cfg:      config.New(),
		log:      logger.New(),
		settings: &settings{},
		mu:       sync.Mutex{},
	}
	flag.StringVar(&app.cfg.Addr, "addr", "localhost:8081", "HTTP network address")
	flag.StringVar(&app.cfg.StaticDir, "static-dir", "./web/static", "Path to static assets")
	flag.Parse()
	return app
}

// Server
func (app *Application) Server(addr string) {
	app.cfg.Addr = addr
	app.srv = server.New(addr)
	app.srv.ErrorLog = app.log.Error

	app.handle()

	app.log.Info.Printf("Listening on port %s", server.DividePortFromAddr(addr))
	app.log.Info.Printf("Open http://%s in the browser", addr)
	log.Fatal(app.srv.Run())
}

func (app *Application) handle() {
	app.srv.HandleFunc("/", handler.Home)
	app.srv.HandleFunc("/about", handler.About)
	app.srv.HandleFunc("/error", handler.Error)
	app.srv.HandleFileServer("./web/static")
}

// Theme
func (app *Application) Theme() string {
	return app.settings.theme
}
