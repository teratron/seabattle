package config

import (
	"path/filepath"
	"time"
)

func init() {
	/*if err := os.Setenv("PORT", "8080"); err != nil {
		log.Println(err)
	}*/
	/*port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}*/
}

type Config struct {
	file string
	addr string

	Server `yaml:"server"`
	Entry  `yaml:"entry"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	Timeout `yaml:"timeout"`
}

type Timeout struct {
	Server time.Duration `yaml:"server"`
	Read   time.Duration `yaml:"read"`
	Write  time.Duration `yaml:"write"`
	Idle   time.Duration `yaml:"idle"`
}

type Entry map[string]map[string]interface{}

func New() *Config {
	cfg := &Config{
		file: filepath.Join("configs", "config.yml"),
		addr: "localhost:8080",
		Server: Server{
			Host: "localhost",
			Port: 8080,
			Timeout: Timeout{
				Server: 30,
				Read:   15,
				Write:  10,
				Idle:   5,
			},
		},
		Entry: map[string]map[string]interface{}{
			"/": {
				"name":  "Home",
				"title": "Sea Battle - Home",
			},
			"/about": {
				"name":  "About",
				"title": "About Us",
			},
		},
	}
	return cfg
}

/*func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./configs/config.yml", "Path to config file")
	//flag.StringVar(&app.cfg.Addr, "addr", "localhost:8081", "HTTP network address")
	//flag.StringVar(&app.cfg.StaticDir, "static-dir", "./web/static", "Path to static assets")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	// Return the configuration path
	return configPath, nil
}

func NewRouter() *http.ServeMux {
	// Create router and define routes and return that router
	router := http.NewServeMux()

	router.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	return router
}

func (cfg Config) Run() {
	// Set up a channel to listen to for interrupt signals
	var runChan = make(chan os.Signal, 1)

	// Set up a context to allow for graceful server shutdowns in the event
	// of an OS interrupt (defers the cancel just in case)
	ctx, cancel := context.WithTimeout(
		context.Background(),
		cfg.Server.Timeout.Server,
	)
	defer cancel()

	// Define server options
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port),
		Handler:      NewRouter(),
		ReadTimeout:  cfg.Server.Timeout.Read * time.Second,
		WriteTimeout: cfg.Server.Timeout.Write * time.Second,
		IdleTimeout:  cfg.Server.Timeout.Idle * time.Second,
	}

	// Handle ctrl+c/ctrl+x interrupt
	signal.Notify(runChan, os.Interrupt, syscall.SIGTRAP)

	// Alert the user that the server is starting
	log.Printf("Server is starting on %s\n", server.Addr)

	// Run the server on a new goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				// Normal interrupt operation, ignore
			} else {
				log.Fatalf("Server failed to start due to err: %v", err)
			}
		}
	}()

	// Block on this channel listeninf for those previously defined syscalls assign
	// to variable so we can let the user know why the server is shutting down
	interrupt := <-runChan

	// If we get one of the pre-prescribed syscalls, gracefully terminate the server
	// while alerting the user
	log.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefully shutdown due to err: %+v", err)
	}
}*/
