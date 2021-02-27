package config

import (
	"path/filepath"
	"time"

	"github.com/teratron/seabattle/pkg/api"
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

	file := api.GetFileType(cfg.file)
	if err, ok := file.(*api.FileError); !ok {
		cfg.Err = cfg.Decode(file)
	} else {
		cfg.Err = err.Err
	}
	return cfg
}

func (cfg *ConfRouter) Decode(decoder api.Decoder) error {
	return decoder.Decode(cfg)
}
