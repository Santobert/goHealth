package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Santobert/gohealth/internal/config"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/load"
)

type Load struct {
	Healthy bool    `json:"healthy"`
	Load1   float64 `json:"load1"`
	Load5   float64 `json:"load5"`
	Load15  float64 `json:"load15"`
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	load, err := load.Avg()
	cpus, cpuErr := cpu.Counts(true)
	if err != nil || cpuErr != nil {
		http.Error(w, "Error retrieving load information", http.StatusInternalServerError)
		return
	}

	healthy := loadHealthy(load.Load1, cpus) && loadHealthy(load.Load5, cpus) && loadHealthy(load.Load15, cpus)
	loadMsg := Load{
		Healthy: healthy,
		Load1:   load.Load1,
		Load5:   load.Load5,
		Load15:  load.Load15,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&loadMsg)
}

func loadHealthy(load float64, cpus int) bool {
	return load/float64(cpus) < config.AppConfig.MaxLoad
}
