package config

type Config struct {
	Version string  `yaml:"version"`
	Module  *Module `yaml:"module"`
}

func NewConfig(modPath string) *Config {
	return &Config{
		Version: "v1alpha",
		Module:  NewModule(modPath),
	}
}
