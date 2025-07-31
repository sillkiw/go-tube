package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the root configuration for the application
// Mode selects "dev" or "production" environment
// Server holds HTTP/TLS settings
// Upload defines upload restrictions
// Video covers conversion and streaming options
// UI holds template/static paths
// Auth controls role-based permissions
type Config struct {
	Mode   string       `yaml:"mode"`
	Server ServerConfig `yaml:"server"`
	Upload UploadConfig `yaml:"upload"`
	Video  VideoConfig  `yaml:"video"`
	UI     UIConfig     `yaml:"ui"`
	Auth   AuthConfig   `yaml:"auth"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
