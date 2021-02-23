package config

import (
	"path/filepath"
	"time"
)

type Server struct {
	File string `yaml:"-"`

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

// New
func (s Server) New() *Server {
	return &Server{
		File: filepath.Join("configs", "server.yml"),
	}
}
