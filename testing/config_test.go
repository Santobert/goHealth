package config

import (
	"reflect"
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
			MaxMemory:   90.0,
			MaxSwap:     90.0,
			SwapEnabled: true,
		},
		Disk: config.DiskConfig{
			MaxDisk: 90.0,
			Paths:   []string{},
			Ignore:  []string{},
			Auto:    true,
		},
		Systemd: config.SystemdConfig{
			Enabled: true,
		},
	}

	if !reflect.DeepEqual(config.AppConfig, expectedConfig) {
		t.Errorf("Expected default config: %+v, got: %+v", expectedConfig, config.AppConfig)
	}
}

func TestReadConfig_FromFile(t *testing.T) {
	config.ReadConfig("config_test.yaml")
	expectedConfig := config.Config{
		Load: config.LoadConfig{
			MaxLoad: 0.9,
		},
		Memory: config.MemoryConfig{
			MaxMemory:   80.0,
			MaxSwap:     70.0,
			SwapEnabled: false,
		},
		Disk: config.DiskConfig{
			MaxDisk: 60.0,
			Paths:   []string{"/", "/home"},
			Ignore:  []string{"/boot/efi"},
			Auto:    false,
		},
		Systemd: config.SystemdConfig{
			Enabled: false,
		},
	}

	if !reflect.DeepEqual(config.AppConfig, expectedConfig) {
		t.Errorf("Expected default config: %+v, got: %+v", expectedConfig, config.AppConfig)
	}
}
