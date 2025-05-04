package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Santobert/gohealth/internal/config"
	"github.com/Santobert/gohealth/internal/handlers"
)

func endpoints(w http.ResponseWriter, r *http.Request) {
	endpoints := []string{
		"/disk",
		"/load",
		"/memory",
		"/config",
		"/systemd",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(endpoints)
}

func main() {
	port := flag.Int("port", 9100, "Server port")
	configFile := flag.String("config", "", "Configuration file")
	flag.Parse()

	config.ReadConfig(*configFile)

	http.HandleFunc("/", endpoints)
	http.HandleFunc("/disk", handlers.DiskUsageHandler)
	http.HandleFunc("/load", handlers.LoadHandler)
	http.HandleFunc("/memory", handlers.MemoryUsageHandler)
	http.HandleFunc("/config", handlers.ConfigHandler)
	http.HandleFunc("/systemd", handlers.SystemdHandler)

	log.Printf("Server is running on port %d...\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
