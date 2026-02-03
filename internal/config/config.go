package config

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the root configuration for the application
// ENV selects "dev" or "prod" environment
// Server holds HTTP/TLS settings
// Upload defines upload restrictions
// Video covers conversion and streaming options
// UI holds template/static paths
// Auth controls role-based permissions
type Config struct {
	Env        string           `yaml:"mode"`
	Server     ServerConfig     `yaml:"server"`
	Upload     UploadConfig     `yaml:"upload"`
	Video      VideoConfig      `yaml:"video"`
	Validation ValidationConfig `yaml:"validation"`
	Auth       AuthConfig       `yaml:"auth"`
}

const configPathDefault = "./config/config.yaml"

func MustLoad() Config {
	configPath := flag.String("config", configPathDefault, "Path to configuration file")
	flag.Parse()

	if _, err := os.Stat(*configPath); err != nil {
		log.Fatalf("config file does not exist: %s", *configPath)
	}

	data, err := os.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return cfg

}
