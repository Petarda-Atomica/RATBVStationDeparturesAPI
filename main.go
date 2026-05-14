package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func departureHandler(w http.ResponseWriter, r *http.Request) {
	// Request payload
	type requestPayload struct {
		Day   int                `json:"day"`
		Lines []LineStationCombo `json:"lines"`
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// Decode incoming JSON
	var payload requestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Received invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call function
	result := getDepartureTimeTable(payload.Day, payload.Lines)

	// Make table
	table := styleTimeTable(result)

	// Send the response back
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(table))
}

func futureDepartureHandler(w http.ResponseWriter, r *http.Request) {
	// Request payload
	type requestPayload struct {
		Count int                `json:"count"`
		Lines []LineStationCombo `json:"lines"`
	}

	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// Decode incoming JSON
	var payload requestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Received invalid request body")
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call function
	now := time.Now()
	result := getDepartureTimeTable((int(now.Weekday()+6) % 7), payload.Lines)

	// Remove missed busses
	nowMin := now.Hour()*60 + now.Minute()
	for i, val := range result {
		if nowMin > val.Hour*60+val.Minute {
			continue
		}
		result = result[i:]
		break
	}

	// Make table
	table := styleTimeTable(result)

	// Send the response back
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(table))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/all-departures", departureHandler)
	mux.HandleFunc("/future-departures", futureDepartureHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
}
