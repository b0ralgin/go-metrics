package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Port   string        `envconfig:"port"`
	File   string        `envconfig:"file"`
	Period time.Duration `envconfig:"period"`
}

// LoadConfig ...
func LoadConfig(app string, c *Config) error {
	if configErr := envconfig.Process(app, c); configErr != nil {
		return configErr
	}
	return nil
}
