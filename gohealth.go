package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Santobert/gohealth/config"
	"github.com/Santobert/gohealth/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := flag.Int("port", 9100, "Server port")
	configFile := flag.String("config", "", "Configuration file")
	flag.Parse()

	config.LoadConfig(*configFile)

	r := mux.NewRouter()
	r.HandleFunc("/disk/{path:.*}", handlers.DiskUsageHandler).Methods("GET")
	r.HandleFunc("/load", handlers.LoadHandler).Methods("GET")
	r.HandleFunc("/memory", handlers.MemoryUsageHandler).Methods("GET")
	r.HandleFunc("/config", handlers.ConfigHandler).Methods("GET")

	log.Printf("Server is running on port %d...\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
