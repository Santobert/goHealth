package config

import (
	"testing"

	"github.com/Santobert/gohealth/internal/config"
)

func TestReadConfig_DefaultValues(t *testing.T) {
	config.ReadConfig("")
	expectedConfig := config.Config{
		Load: config.LoadConfig{
			MaxLoad: 1.0,
		},
		Memory: config.MemoryConfig{
			MaxMemory: 90.0,
			MaxSwap:   90.0,
		},
		Disk: config.DiskConfig{
			MaxDisk: 90.0,
			Paths:   []string{"/"},
		},
	}

	assertDefaultConfig(t, expectedConfig)
}

func TestReadConfig_FromFile(t *testing.T) {
	config.ReadConfig("config_test.yaml")
	expectedConfig := config.Config{
		Load: config.LoadConfig{
			MaxLoad: 0.9,
		},
		Memory: config.MemoryConfig{
			MaxMemory: 80.0,
			MaxSwap:   70.0,
		},
		Disk: config.DiskConfig{
			MaxDisk: 60.0,
			Paths:   []string{"/", "/home"},
		},
	}

	assertDefaultConfig(t, expectedConfig)
}

func assertDefaultConfig(t *testing.T, expectedConfig config.Config) {
	if config.AppConfig.Load != expectedConfig.Load ||
		config.AppConfig.Memory != expectedConfig.Memory ||
		config.AppConfig.Disk.MaxDisk != expectedConfig.Disk.MaxDisk ||
		len(config.AppConfig.Disk.Paths) != len(expectedConfig.Disk.Paths) {
		t.Errorf("Expected config to be %+v, got %+v", expectedConfig, config.AppConfig)
	}
	for i, path := range expectedConfig.Disk.Paths {
		if config.AppConfig.Disk.Paths[i] != path {
			t.Errorf("Expected Paths[%d] to be %s, got %s", i, path, config.AppConfig.Disk.Paths[i])
		}
	}
}
