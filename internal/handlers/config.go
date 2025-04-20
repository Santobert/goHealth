package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Santobert/gohealth/internal/config"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config.AppConfig)
}
