/*
Package config provides config structures and parse funcs.
*/
package config

// Config is the global configuration.
type Config struct {
	App    ApplicationConfig   `json:"app" yaml:"app"`
	DB     DBConfig            `json:"db" yaml:"db"`
	MQ     MQConfig            `json:"mq" yaml:"mq"`
	Remote RemoteServiceConfig `json:"remote" yaml:"remote"`
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

// MQConfig is the configuration of message queues.
type MQConfig struct {
	Brokers []string `json:"brokers" yaml:"brokers"`
	Topic   string   `json:"topic" yaml:"topic"`
}

// RemoteServiceConfig is the configuration of remote services.
type RemoteServiceConfig struct {
	TrendURL string `json:"trend_url" yaml:"trend_url"`
	RecURL   string `json:"rec_url" yaml:"rec_url"`
}
