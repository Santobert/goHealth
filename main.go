package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Santobert/gohealth/internal/config"
	"github.com/Santobert/gohealth/internal/handlers"
)

func main() {
	port := flag.Int("port", 9100, "Server port")
	configFile := flag.String("config", "", "Configuration file")
	flag.Parse()

	config.ReadConfig(*configFile)

	http.HandleFunc("/disk", handlers.DiskUsageHandler)
	http.HandleFunc("/load", handlers.LoadHandler)
	http.HandleFunc("/memory", handlers.MemoryUsageHandler)
	http.HandleFunc("/config", handlers.ConfigHandler)

	log.Printf("Server is running on port %d...\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
