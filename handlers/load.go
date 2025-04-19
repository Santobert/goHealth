package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/shirou/gopsutil/v4/load"
)

type Load struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

func getLoad() (*Load, error) {
	load, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return &Load{
		Load1:  load.Load1,
		Load5:  load.Load5,
		Load15: load.Load15,
	}, nil
}

func LoadHandler(w http.ResponseWriter, r *http.Request) {
	load, err := getLoad()
	if err != nil {
		http.Error(w, "Error retrieving load information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(load)
}
