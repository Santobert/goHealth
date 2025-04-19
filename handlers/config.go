package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Santobert/gohealth/config"
)

type ConfigResponse struct {
	MaxLoad   float64 `json:"max_load"`
	MaxMemory float64 `json:"max_memory"`
	MaxSwap   float64 `json:"max_swap"`
	MaxDisk   float64 `json:"max_disk"`
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	response := ConfigResponse{
		MaxLoad:   config.AppConfig.MaxLoad,
		MaxMemory: config.AppConfig.MaxMemory,
		MaxSwap:   config.AppConfig.MaxSwap,
		MaxDisk:   config.AppConfig.MaxDisk,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
