package user

import (
	"os"

	"gopkg.in/yaml.v3"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Role     string `yaml:"role"`
}

// LoadUsersFromFile loads users from a YAML file and stores them in the global Users slice.
func LoadUsersFromFile(path string) ([]User, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var loaded []User
	if err := yaml.Unmarshal(data, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}
