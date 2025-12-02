package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type breakRecommendationRequest struct {
	TotalSessions int `json:"total_sessions"`
}

type breakRecommendationResponse struct {
	TakeBreak bool `json:"take_break"`
}

func recommendBreak(w http.ResponseWriter, req *http.Request) {
	// Only allow POST requests
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	defer req.Body.Close()
	var request breakRecommendationRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	takeBreak := false
	if (request.TotalSessions%3 == 0) && request.TotalSessions > 0 {
		takeBreak = true
	}

	response := breakRecommendationResponse{
		TakeBreak: takeBreak,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/break", recommendBreak)

	fmt.Println("Break Recommendation Microservice running on http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
