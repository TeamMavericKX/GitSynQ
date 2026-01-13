package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Project ProjectConfig `yaml:"project"`
	Server  ServerConfig  `yaml:"server"`
	Bundle  BundleConfig  `yaml:"bundle"`
}

type ProjectConfig struct {
	Name   string `yaml:"name"`
	Branch string `yaml:"branch"`
}

type ServerConfig struct {
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	Port       int    `yaml:"port"`
	RemotePath string `yaml:"remote_path"`
	SSHKeyPath string `yaml:"ssh_key_path,omitempty"`
}

type BundleConfig struct {
	Directory  string `yaml:"directory"`
	Compress   bool   `yaml:"compress"`
	MaxHistory int    `yaml:"max_history"`
}

const ConfigFile = ".gitsync.yaml"

func Load() (*Config, error) {
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Set defaults
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 22
	}
	if cfg.Bundle.Directory == "" {
		cfg.Bundle.Directory = ".gitsync-bundles"
	}
	if cfg.Bundle.MaxHistory == 0 {
		cfg.Bundle.MaxHistory = 10
	}

	return &cfg, nil
}

func Save(cfg Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(ConfigFile, data, 0644)
}
