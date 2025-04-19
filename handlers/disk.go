package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/shirou/gopsutil/v4/disk"
)

type DiskUsage struct {
	Healthy     bool    `json:"healthy"`
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
		Healthy:     usage.UsedPercent < 90,
		Total:       usage.Total,
		Used:        usage.Used,
		Free:        usage.Free,
		UsedPercent: usage.UsedPercent,
	}, nil
}

func DiskUsageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["path"] == "" {
		http.Redirect(w, r, "/disk/_", http.StatusMovedPermanently)
		return
	}

	path := strings.ReplaceAll(vars["path"], "_", "/")
	usage, err := getDiskUsage(path)
	if err != nil {
		http.Error(w, "Error retrieving disk usage information", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usage)
}
