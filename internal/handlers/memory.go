package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Santobert/gohealth/internal/config"
	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryUsage struct {
	Healthy      bool    `json:"healthy"`
	UsagePercent float64 `json:"usage_percent"`
	SwapPercent  float64 `json:"swap_percent"`
}

func isMemoryHealthy(value float64) bool {
	return value < config.AppConfig.Memory.MaxMemory
}

func isSwapHealthy(value float64) bool {
	return value < config.AppConfig.Memory.MaxSwap
}

func isSwapEnabled() bool {
	return config.AppConfig.Memory.SwapEnabled
}

func addSwapHealth(memoryUsage *MemoryUsage) error {
	swapUsage, swapErr := mem.SwapMemory()
	if swapErr != nil {
		return swapErr
	}
	swapHealthy := isSwapHealthy(swapUsage.UsedPercent)
	memoryUsage.SwapPercent = swapUsage.UsedPercent
	memoryUsage.Healthy = memoryUsage.Healthy && swapHealthy
	return nil
}

func MemoryUsageHandler(w http.ResponseWriter, r *http.Request) {
	memoryUsage, memErr := mem.VirtualMemory()
	if memErr != nil {
		http.Error(w, "Failed to retrieve memory usage information", http.StatusInternalServerError)
		return
	}

	memHealthy := isMemoryHealthy(memoryUsage.UsedPercent)
	memMsg := MemoryUsage{
		Healthy:      memHealthy,
		UsagePercent: memoryUsage.UsedPercent,
		SwapPercent:  0,
	}

	if isSwapEnabled() {
		if err := addSwapHealth(&memMsg); err != nil {
			http.Error(w, "Failed to retrieve swap memory usage information", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&memMsg)
}
