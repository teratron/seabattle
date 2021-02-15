package config

import (
	"fmt"
)

// Config is global object that holds all application level variables.
//var zConfig Config

type Config struct {
	// Example Variable
	//ConfigVar string

	Addr      string
	StaticDir string
}

func New() *Config {
	return &Config{}
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	/*v := viper.New()
	v.SetConfigName("example")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)*/
	return fmt.Errorf("failed to read the configuration file")
}
