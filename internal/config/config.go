package config

import (
	"bytes"
	"fmt"

	"github.com/hjblom/fuse/internal/util"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Version string  `yaml:"version"`
	Module  *Module `yaml:"module"`
}

func NewConfig(modPath string) *Config {
	return &Config{
		Module: NewModule(modPath),
	}
}

func ReadConfig(path string, fi util.FileInterface) (*Config, error) {
	// Read file
	data, err := fi.Read(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal
	c := &Config{}
	err = yaml.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return c, nil
}

func WriteConfig(c *Config, path string, fi util.FileInterface) error {
	// Setup YAML encoder
	data := &bytes.Buffer{}
	enc := yaml.NewEncoder(data)
	defer enc.Close()
	enc.SetIndent(2)

	// Marshal
	err := enc.Encode(c)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write file
	err = fi.Write(path, data.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
