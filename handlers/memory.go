package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shirou/gopsutil/v4/mem"
)

type MemoryUsage struct {
	UsagePercent float64 `json:"usage_percent"`
	SwapPercent  float64 `json:"swap_percent"`
}

func getMemoryUsage() (*MemoryUsage, error) {
	memory, memErr := mem.VirtualMemory()
	swap, swapErr := mem.SwapMemory()

	if memErr != nil || swapErr != nil {
		return nil, fmt.Errorf("memory error: %v, swap error: %v", memErr, swapErr)
	}

	return &MemoryUsage{
		UsagePercent: memory.UsedPercent,
		SwapPercent:  swap.UsedPercent,
	}, nil
}

func MemoryUsageHandler(w http.ResponseWriter, r *http.Request) {
	memoryUsage, err := getMemoryUsage()
	if err != nil {
		http.Error(w, "Failed to retrieve memory usage information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(memoryUsage)
}
