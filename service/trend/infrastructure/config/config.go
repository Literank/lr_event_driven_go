/*
Package config provides config structures and parse funcs.
*/
package config

// Config is the global configuration.
type Config struct {
	App   ApplicationConfig `json:"app" yaml:"app"`
	Cache CacheConfig       `json:"cache" yaml:"cache"`
}

// ApplicationConfig is the configuration of main app.
type ApplicationConfig struct {
	Port int `json:"port" yaml:"port"`
}

// CacheConfig is the configuration of cache.
type CacheConfig struct {
	Address  string `json:"address" yaml:"address"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}
