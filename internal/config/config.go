// Package config handles GitSynq configuration loading and saving.
package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the main configuration structure for GitSynq.
type Config struct {
	Project ProjectConfig `yaml:"project"`
	Server  ServerConfig  `yaml:"server"`
	Bundle  BundleConfig  `yaml:"bundle"`
}

// ProjectConfig contains details about the local Git repository.
type ProjectConfig struct {
	Name   string `yaml:"name"`
	Branch string `yaml:"branch"`
}

// ServerConfig contains connection details for the remote air-gapped server.
type ServerConfig struct {
	Host       string `yaml:"host"`
	User       string `yaml:"user"`
	Port       int    `yaml:"port"`
	RemotePath string `yaml:"remote_path"`
	SSHKeyPath string `yaml:"ssh_key_path,omitempty"`
}

// BundleConfig contains settings for Git bundle creation and storage.
type BundleConfig struct {
	Directory  string `yaml:"directory"`
	Compress   bool   `yaml:"compress"`
	MaxHistory int    `yaml:"max_history"`
}

// ConfigFile is the default name for the GitSynq configuration file.
var ConfigFile = ".gitsync.yaml"

// Load reads and parses the configuration file from the current directory.
func Load() (*Config, error) {
	data, err := os.ReadFile(ConfigFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
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

// Save marshals and writes the configuration to the local disk.
func Save(cfg Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(ConfigFile, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
