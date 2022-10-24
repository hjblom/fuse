package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Module     string      `yaml:"module"`
	Components []Component `yaml:"components"`
}

type Component struct {
	Package  string   `yaml:"package"`
	Path     string   `yaml:"path,omitempty"`
	Tags     []string `yaml:"tags,omitempty"`
	Requires []string `yaml:"requires,omitempty"`
}

func LoadConfigFile(path string) (*Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	c, err := parseConfigYAML(b)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return c, nil
}

func WriteConfigFile(path string, c *Config) error {
	b, err := marshalConfigYAML(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	return nil
}

func parseConfigYAML(b []byte) (*Config, error) {
	c := &Config{}
	err := yaml.Unmarshal(b, c)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return c, nil
}

func marshalConfigYAML(c *Config) ([]byte, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config: %w", err)
	}
	return b, nil
}
