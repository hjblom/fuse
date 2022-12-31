package config

type PackageConfig struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Env         string `yaml:"env"`
	Required    bool   `yaml:"required"`
}
