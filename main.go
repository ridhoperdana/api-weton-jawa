package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var indonesianWeekdays = map[string]string{
	"Sunday":    "Minggu",
	"Monday":    "Senin",
	"Tuesday":   "Selasa",
	"Wednesday": "Rabu",
	"Thursday":  "Kamis",
	"Friday":    "Jumat",
	"Saturday":  "Sabtu",
}

var pasaran = []string{"Legi", "Pahing", "Pon", "Wage", "Kliwon"}

// GetPasaran calculates the Javanese "pasaran" day for a given Gregorian date.
func GetPasaran(t time.Time) string {
	// Reference date: 1 January 1893 was a "Kliwon"
	referenceDate := time.Date(1893, 1, 1, 0, 0, 0, 0, time.UTC)
	days := int(t.Sub(referenceDate).Hours() / 24)

	// Calculate pasaran by modulus 5 since there are 5 pasaran days
	pasaranIndex := (days%5 + 5) % 5 // Handles negative modulus result
	return pasaran[pasaranIndex]
}

// WetonResponse represents the structure of the response
type WetonResponse struct {
	Data     WetonData `json:"data"`
	Metadata struct{}  `json:"metadata"`
}

// WetonData holds the day and pasaran information
type WetonData struct {
	Hari    string `json:"hari"`
	Pasaran string `json:"pasaran"`
}

// GetWetonHandler handles the /weton/{tanggal-masehi} endpoint
func GetWetonHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the date from the URL path
	tanggal := strings.TrimPrefix(r.URL.Path, "/api/weton/")
	tanggal = strings.TrimSpace(tanggal)

	// Parse the input date in dd-mm-yyyy format
	layout := "02-01-2006"
	date, err := time.Parse(layout, tanggal)
	if err != nil {
		http.Error(w, "Invalid date format. Use DD-MM-YYYY.", http.StatusBadRequest)
		return
	}

	// Get the day in Indonesian
	gregorianDay := indonesianWeekdays[date.Weekday().String()]

	// Get the Pasaran day
	pasaranDay := GetPasaran(date)

	// Create the response
	response := WetonResponse{
		Data: WetonData{
			Hari:    gregorianDay,
			Pasaran: pasaranDay,
		},
	}

	// Set response headers and return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Define the route handler
	http.HandleFunc("/api/weton/", GetWetonHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
