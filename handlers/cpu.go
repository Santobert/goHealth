package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

type CPUUsage struct {
	UsagePercent []float64 `json:"usage_percent"`
}

func getCpuUsage() ([]float64, error) {
	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	return percentages, nil
}

func CpuUsageHandler(w http.ResponseWriter, r *http.Request) {
	percentages, err := getCpuUsage()
	if err != nil {
		http.Error(w, "Fehler beim Abrufen der CPU-Auslastung", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(percentages)
}
