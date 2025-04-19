package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	MaxLoad   float64 `yaml:"max_load,omitempty"`
	MaxMemory float64 `yaml:"max_memory,omitempty"`
	MaxSwap   float64 `yaml:"max_swap,omitempty"`
	MaxDisk   float64 `yaml:"max_disk,omitempty"`
}

var AppConfig Config

func LoadConfig(filename string) {
	AppConfig = Config{
		MaxLoad:   1,
		MaxMemory: 90,
		MaxSwap:   90,
		MaxDisk:   90,
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
