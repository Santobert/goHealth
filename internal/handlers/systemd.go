package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Santobert/gohealth/internal/config"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/coreos/go-systemd/v22/util"
)

type Systemd struct {
	Healthy     bool     `json:"healthy"`
	FailedUnits []string `json:"failed_units"`
}

func getFailedUnits(ctx context.Context) ([]string, error) {
	conn, err := dbus.NewSystemConnectionContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	units, err := conn.ListUnitsFilteredContext(ctx, []string{"failed"})
	if err != nil {
		return nil, err
	}

	var failedUnits []string
	for _, unit := range units {
		failedUnits = append(failedUnits, unit.Name)
	}
	return failedUnits, nil
}

func SystemdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	if !config.AppConfig.Systemd.Enabled || !util.IsRunningSystemd() {
		http.Error(w, "Systemd monitoring is disabled or systemd is not running", http.StatusNotFound)
		return
	}

	failedUnits, err := getFailedUnits(ctx)
	if err != nil {
		http.Error(w, "Failed to get systemd units", http.StatusInternalServerError)
		return
	}
	systemdMsg := Systemd{
		Healthy:     len(failedUnits) == 0,
		FailedUnits: failedUnits,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(systemdMsg)
}
