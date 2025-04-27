package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type LoadConfig struct {
	MaxLoad float64 `yaml:"max_load,omitempty"`
}

type MemoryConfig struct {
	MaxMemory   float64 `yaml:"max_memory,omitempty"`
	MaxSwap     float64 `yaml:"max_swap,omitempty"`
	SwapEnabled bool    `yaml:"swap_enabled,omitempty"`
}

type DiskConfig struct {
	MaxDisk float64  `yaml:"max_disk,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
	Ignore  []string `yaml:"ignore,omitempty"`
	Auto    bool     `yaml:"auto,omitempty"`
}

type Config struct {
	Load   LoadConfig   `yaml:"load,omitempty"`
	Memory MemoryConfig `yaml:"memory,omitempty"`
	Disk   DiskConfig   `yaml:"disk,omitempty"`
}

var AppConfig Config

func ReadConfig(filename string) {
	AppConfig = Config{
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
	}

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
