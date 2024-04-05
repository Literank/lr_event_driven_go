/*
Package config provides config structures and parse funcs.
*/
package config

// Config is the global configuration.
type Config struct {
	App ApplicationConfig `json:"app" yaml:"app"`
	DB  DBConfig          `json:"db" yaml:"db"`
	MQ  MQConfig          `json:"mq" yaml:"mq"`
}

// DBConfig is the configuration of databases.
type DBConfig struct {
	MongoURI    string `json:"mongo_uri" yaml:"mongo_uri"`
	MongoDBName string `json:"mongo_db_name" yaml:"mongo_db_name"`
}

// ApplicationConfig is the configuration of main app.
type ApplicationConfig struct {
	Port     int `json:"port" yaml:"port"`
	PageSize int `json:"page_size" yaml:"page_size"`
}

// MQConfig is the configuration of message queues.
type MQConfig struct {
	Brokers []string `json:"brokers" yaml:"brokers"`
	Topic   string   `json:"topic" yaml:"topic"`
	GroupID string   `json:"group_id" yaml:"group_id"`
}
