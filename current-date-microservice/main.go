package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type dateResponse struct {
	Date string `json:"date"`
}

func currentDate(w http.ResponseWriter, req *http.Request) {
	// Only allow GET requests
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the current date
	current := time.Now().Format("01/02/2006")

	response := dateResponse {
		Date: current,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/date", currentDate)

	fmt.Println("Current Date Microservice running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}