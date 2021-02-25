package config

import (
	"path/filepath"
	"time"
)

type ConfServer struct {
	file string
	Err  error `yaml:"-"`

	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Timeout `yaml:"timeout"`
}

type Timeout struct {
	Header time.Duration `yaml:"header"`
	Read   time.Duration `yaml:"read"`
	Write  time.Duration `yaml:"write"`
	Idle   time.Duration `yaml:"idle"`
}

// New
/*func (cfg *Server) New() *Server {
	cfg.file = filepath.Join("configs", "server.yml")
	if cfg.Err = cfg.Decode(cfg.file); cfg.Err != nil {
		cfg.Host = "localhost"
		cfg.Port = 8080
		cfg.Header = 30
		cfg.Read = 15
		cfg.Write = 10
		cfg.Idle = 5
	}
	return cfg
}*/

func NewConfServer() *ConfServer {
	cfg := &ConfServer{
		file: filepath.Join("configs", "server.yml"),
		Host: "localhost",
		Port: 8080,
		Timeout: Timeout{
			Header: 30,
			Read:   15,
			Write:  10,
			Idle:   5,
		},
	}
	cfg.Err = cfg.Decode(cfg.file)
	return cfg
}
