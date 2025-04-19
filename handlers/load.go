package handlers

import (
	"encoding/json"
	"net/http"

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

	loadMsg := Load{
		Healthy: load.Load1 < float64(cpus) && load.Load5 < float64(cpus) && load.Load15 < float64(cpus),
		Load1:   load.Load1,
		Load5:   load.Load5,
		Load15:  load.Load15,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&loadMsg)
}
