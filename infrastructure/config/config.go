/*
Package config provides config structures and parse funcs.
*/
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the global configuration.
type Config struct {
	App ApplicationConfig `json:"app" yaml:"app"`
	DB  DBConfig          `json:"db" yaml:"db"`
}

// DBConfig is the configuration of databases.
type DBConfig struct {
	DSN string `json:"dsn" yaml:"dsn"`
}

// ApplicationConfig is the configuration of main app.
type ApplicationConfig struct {
	Port             int    `json:"port" yaml:"port"`
	PageSize         int    `json:"page_size" yaml:"page_size"`
	TemplatesPattern string `json:"templates_pattern" yaml:"templates_pattern"`
}

// Parse parses config file and returns a Config.
func Parse(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file %s: %v", filename, err)
	}
	return c, nil
}
