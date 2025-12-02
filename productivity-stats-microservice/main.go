package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	totalSessions int
	mu            sync.Mutex
)

type productivityStatsResponse struct {
	TotalSessions int `json:"total_sessions"`
}

func setProductivityStats(w http.ResponseWriter, req *http.Request) {
	// Only allow POST requests
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mu.Lock()
	// Increment counter
	totalSessions += 1
	current := totalSessions
	mu.Unlock()

	response := productivityStatsResponse{
		TotalSessions: current,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Reset stats
	totalSessions = 0
	http.HandleFunc("/stats", setProductivityStats)

	fmt.Println("Productivity Stats Microservice running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
