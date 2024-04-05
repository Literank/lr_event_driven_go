/*
Package config provides config structures and parse funcs.
*/
package config

// Config is the global configuration.
type Config struct {
	App   ApplicationConfig `json:"app" yaml:"app"`
	Cache CacheConfig       `json:"cache" yaml:"cache"`
	MQ    MQConfig          `json:"mq" yaml:"mq"`
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

// MQConfig is the configuration of message queues.
type MQConfig struct {
	Brokers []string `json:"brokers" yaml:"brokers"`
	Topic   string   `json:"topic" yaml:"topic"`
	GroupID string   `json:"group_id" yaml:"group_id"`
}
