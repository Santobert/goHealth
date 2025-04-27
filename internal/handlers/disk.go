package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"

	"github.com/Santobert/gohealth/internal/config"
	"github.com/shirou/gopsutil/v4/disk"
)

type Partition struct {
	Healthy     bool    `json:"-"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskUsage struct {
	Healthy bool                  `json:"healthy"`
	Paths   map[string]*Partition `json:"paths"`
}

func getPartition(path string) (*Partition, error) {
	usage, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}

	healthy := usage.UsedPercent < config.AppConfig.Disk.MaxDisk
	return &Partition{
		Healthy:     healthy,
		Total:       usage.Total,
		Used:        usage.Used,
		Free:        usage.Free,
		UsedPercent: usage.UsedPercent,
	}, nil
}

func addPartitionUsage(diskUsage *DiskUsage, path string) error {
	usage, err := getPartition(path)
	if err != nil {
		return err
	}
	diskUsage.Paths[path] = usage
	diskUsage.Healthy = diskUsage.Healthy && usage.Healthy
	return nil
}

func getAllPartitions() ([]string, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	var partitionsList []string
	for _, partition := range partitions {
		if !slices.Contains(config.AppConfig.Disk.Ignore, partition.Mountpoint) {
			partitionsList = append(partitionsList, partition.Mountpoint)
		}
	}
	return partitionsList, nil

}

func DiskUsageHandler(w http.ResponseWriter, r *http.Request) {
	diskUsage := &DiskUsage{
		Healthy: true,
		Paths:   make(map[string]*Partition),
	}

	paths := []string{}
	if config.AppConfig.Disk.Auto {
		if autoPaths, err := getAllPartitions(); err != nil {
			log.Printf("Error retrieving partitions: %v", err)
		} else {
			paths = autoPaths
		}
	} else {
		paths = config.AppConfig.Disk.Paths
	}

	for _, path := range paths {
		if err := addPartitionUsage(diskUsage, path); err != nil {
			log.Printf("Error retrieving disk usage for path %s: %v", path, err)
			diskUsage.Paths[path] = nil
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&diskUsage)
}
