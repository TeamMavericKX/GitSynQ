package config

import (
	"os"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	tmpFile := ".gitsync.test.yaml"
	// Override ConfigFile for testing
	oldConfigFile := ConfigFile
	defer func() { ConfigFile = oldConfigFile }()
	ConfigFile = tmpFile
	defer os.ReadFile(tmpFile)
	defer os.Remove(tmpFile)

	cfg := Config{
		Project: ProjectConfig{Name: "test-project", Branch: "main"},
		Server:  ServerConfig{Host: "1.2.3.4", User: "tester"},
		Bundle:  BundleConfig{Directory: ".bundles", MaxHistory: 5},
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save() failed: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	if loaded.Project.Name != "test-project" {
		t.Errorf("Expected project name test-project, got %s", loaded.Project.Name)
	}
	if loaded.Server.Host != "1.2.3.4" {
		t.Errorf("Expected host 1.2.3.4, got %s", loaded.Server.Host)
	}
	if loaded.Bundle.MaxHistory != 5 {
		t.Errorf("Expected max history 5, got %d", loaded.Bundle.MaxHistory)
	}
}

func TestLoadDefaults(t *testing.T) {
	tmpFile := ".gitsync.defaults.yaml"
	oldConfigFile := ConfigFile
	defer func() { ConfigFile = oldConfigFile }()
	ConfigFile = tmpFile
	defer os.Remove(tmpFile)

	data := []byte(`
project:
  name: test
server:
  host: localhost
`)
	if err := os.WriteFile(tmpFile, data, 0644); err != nil {
		t.Fatal(err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	if loaded.Server.Port != 22 {
		t.Errorf("Expected default port 22, got %d", loaded.Server.Port)
	}
	if loaded.Bundle.Directory != ".gitsync-bundles" {
		t.Errorf("Expected default directory .gitsync-bundles, got %s", loaded.Bundle.Directory)
	}
}
