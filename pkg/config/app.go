package config

type App struct {
	file string

	Application string `yaml:"application"`
	Version     string `yaml:"version"`
	Runtime     string `yaml:"runtime"`
	ApiVersion  string `yaml:"api_version"`
}
