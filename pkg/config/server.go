package config

import "time"

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
