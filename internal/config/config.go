package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type LoadConfig struct {
	MaxLoad float64 `yaml:"max_load"`
}

type MemoryConfig struct {
	MaxMemory   float64 `yaml:"max_memory"`
	MaxSwap     float64 `yaml:"max_swap"`
	SwapEnabled bool    `yaml:"swap_enabled"`
}

type DiskConfig struct {
	MaxDisk float64  `yaml:"max_disk"`
	Paths   []string `yaml:"paths"`
	Ignore  []string `yaml:"ignore"`
	Auto    bool     `yaml:"auto"`
}

type SystemdConfig struct {
	Enabled bool `yaml:"enabled"`
}

type Config struct {
	Load    LoadConfig    `yaml:"load"`
	Memory  MemoryConfig  `yaml:"memory"`
	Disk    DiskConfig    `yaml:"disk"`
	Systemd SystemdConfig `yaml:"systemd"`
}

var AppConfig Config

var defaultConfig = Config{
	Load: LoadConfig{
		MaxLoad: 1.0,
	},
	Memory: MemoryConfig{
		MaxMemory:   90.0,
		MaxSwap:     90.0,
		SwapEnabled: true,
	},
	Disk: DiskConfig{
		MaxDisk: 90.0,
		Paths:   []string{},
		Ignore:  []string{},
		Auto:    true,
	},
	Systemd: SystemdConfig{
		Enabled: true,
	},
}

func ReadConfig(filename string) {
	AppConfig = defaultConfig

	if filename == "" {
		log.Println("No configuration file provided, using default values.")
		return
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error loading configuration file: %v", err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %v", err)
	}
}
