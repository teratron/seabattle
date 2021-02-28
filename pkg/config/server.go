package config

import (
	"path/filepath"
	"time"

	"github.com/teratron/seabattle/pkg/api"
)

type ConfServer struct {
	file string
	Err  error `json:"-" yaml:"-"`

	Host    string `json:"host" yaml:"host"`
	Port    int    `json:"port" yaml:"port"`
	Timeout `json:"timeout" yaml:"timeout"`
}

type Timeout struct {
	Header time.Duration `json:"header" yaml:"header"`
	Read   time.Duration `json:"read" yaml:"read"`
	Write  time.Duration `json:"write" yaml:"write"`
	Idle   time.Duration `json:"idle" yaml:"idle"`
}

// NewConfServer
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

	file := api.GetFileType(cfg.file)
	if err, ok := file.(*api.FileError); !ok {
		cfg.Err = cfg.Decode(file)
	} else {
		cfg.Err = err.Err
	}
	return cfg
}

func (cfg *ConfServer) Decode(decoder api.Decoder) error {
	return decoder.Decode(cfg)
}
