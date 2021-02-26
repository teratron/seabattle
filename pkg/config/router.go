package config

import (
	"path/filepath"
	"time"
)

type ConfRouter struct {
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

// NewConfServer
func NewConfRouter() *ConfRouter {
	cfg := &ConfRouter{
		file: filepath.Join("configs", "router.yml"),
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
