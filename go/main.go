package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Reading struct {
	ID          int     `json:"id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Pressure    float64 `json:"pressure"`
	Timestamp   string  `json:"timestamp"`
}

func main() {
	// Connect to SQLite database
	db, err := sql.Open("sqlite3", "./python/sensor.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// HTTP handler to retrieve and serve readings as JSON
	http.HandleFunc("/readings", func(w http.ResponseWriter, r *http.Request) {
		// Query the database for all environmental readings
		rows, err := db.Query("SELECT id, temperature, humidity, pressure, timestamp FROM environmental_readings")
		if err != nil {
			http.Error(w, "Failed to query database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var readings []Reading

		// Loop through the result set and populate the readings slice
		for rows.Next() {
			var reading Reading
			err := rows.Scan(&reading.ID, &reading.Temperature, &reading.Humidity, &reading.Pressure, &reading.Timestamp)
			if err != nil {
				http.Error(w, "Failed to scan row", http.StatusInternalServerError)
				return
			}
			readings = append(readings, reading)
		}

		// Convert readings to JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(readings)
		if err != nil {
			http.Error(w, "Failed to encode readings as JSON", http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
