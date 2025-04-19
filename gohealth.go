package main

import (
	"fmt"
	"net/http"

	"github.com/Santobert/gohealth/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/disk/?", handlers.DiskUsageHandler).Methods("GET")
	r.HandleFunc("/disk/{path}", handlers.DiskUsageHandler).Methods("GET")
	r.HandleFunc("/cpu", handlers.CpuUsageHandler).Methods("GET")
	r.HandleFunc("/load", handlers.LoadHandler).Methods("GET")
	r.HandleFunc("/memory", handlers.MemoryUsageHandler).Methods("GET")

	fmt.Println("Server läuft auf Port 8080...")
	http.ListenAndServe(":8080", r)
}
