package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func diskUsageHandler(w http.ResponseWriter, r *http.Request) {
	usage, err := getDiskUsage("/")
	if err != nil {
		http.Error(w, "Fehler beim Abrufen der Festplattenauslastung", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usage)
}

func main() {
	http.HandleFunc("/disk", diskUsageHandler)
	fmt.Println("Server läuft auf Port 8080...")
	http.ListenAndServe(":8080", nil)
}
