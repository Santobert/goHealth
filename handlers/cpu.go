package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CpuUsage struct {
	UsagePercent float64 `json:"usage_percent"`
}

func getCpuUsage() (*CpuUsage, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	return &CpuUsage{
		UsagePercent: percentages[0],
	}, nil
}

func CpuUsageHandler(w http.ResponseWriter, r *http.Request) {
	cpu, err := getCpuUsage()
	if err != nil {
		http.Error(w, "Error retrieving CPU usage information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cpu)
}
