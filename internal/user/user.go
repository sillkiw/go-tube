package user

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}

// MustLoad loads users from a YAML file and stores them in the global Users slice.
func MustLoad(path string) []User {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read user files: %v", err)
	}

	var loaded []User
	if err := yaml.Unmarshal(data, &loaded); err != nil {
		log.Fatalf("failed to unmarshal user file: %v", err)
	}

	return loaded
}
