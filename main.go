package main

import (
	"fmt"
	"time"
)

// Map of weekdays in English to Indonesian
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

// IsWithinRange checks if a date is within the range of 1990-2000.
func IsWithinRange(t time.Time) bool {
	startDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2000, 12, 31, 23, 59, 59, 0, time.UTC)
	return t.After(startDate) && t.Before(endDate)
}

func main() {
	// Input the date you want to convert in DD-MM-YYYY format
	inputDate := "30-04-1995" // Example: 30 April 1995
	layout := "02-01-2006"
	date, err := time.Parse(layout, inputDate)

	if err != nil {
		fmt.Println("Invalid date format. Use DD-MM-YYYY.")
		return
	}

	// Check if the date is within 1990-2000
	if !IsWithinRange(date) {
		fmt.Println("Date is out of range. Please input a date between 1990 and 2000.")
		return
	}

	// Get the Pasaran day
	pasaranDay := GetPasaran(date)

	// Get the day of the week in Indonesian
	gregorianDay := indonesianWeekdays[date.Weekday().String()]

	// Output the Javanese calendar conversion in DD-MM-YYYY format
	fmt.Printf("Hari jawa kamu adalah: %s %s\n", gregorianDay, pasaranDay)
}
