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

func MemoryUsageHandler(w http.ResponseWriter, r *http.Request) {
	memoryUsage, memErr := mem.VirtualMemory()
	swapUsage, swapErr := mem.SwapMemory()
	if memErr != nil || swapErr != nil {
		http.Error(w, "Failed to retrieve memory usage information", http.StatusInternalServerError)
		return
	}

	memHealthy := isMemoryHealthy(memoryUsage.UsedPercent)
	swapHealty := isSwapHealthy(swapUsage.UsedPercent)
	memMsg := MemoryUsage{
		Healthy:      memHealthy && swapHealty,
		UsagePercent: memoryUsage.UsedPercent,
		SwapPercent:  swapUsage.UsedPercent,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&memMsg)
}
