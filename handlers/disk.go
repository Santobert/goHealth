package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

func getDiskUsage(path string) (*DiskUsage, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}

	return &DiskUsage{
		Total:       usage.Total,
		Used:        usage.Used,
		Free:        usage.Free,
		UsedPercent: usage.UsedPercent,
	}, nil
}

func DiskUsageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	path := strings.ReplaceAll(vars["path"], "_", "/")

	usage, err := getDiskUsage(path)
	if err != nil {
		http.Error(w, "Fehler beim Abrufen der Festplattenauslastung", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usage)
}
