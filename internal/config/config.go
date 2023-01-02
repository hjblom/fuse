package config

const DefaultVersion = "v1alpha"

type Config struct {
	Version string  `yaml:"version"`
	Module  *Module `yaml:"module"`
}

type ConfigOption func(*Config)

func WithModulePath(path string) ConfigOption {
	return func(m *Config) {
		m.Module.Path = path
	}
}

func NewConfig(opts ...ConfigOption) *Config {
	c := &Config{
		Version: DefaultVersion,
		Module:  &Module{},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
