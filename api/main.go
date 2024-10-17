package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

var neptuHari = map[string]int{
	"Minggu": 5,
	"Senin":  4,
	"Selasa": 3,
	"Rabu":   7,
	"Kamis":  8,
	"Jumat":  6,
	"Sabtu":  9,
}

var neptuPasaran = map[string]int{
	"Legi":   5,
	"Pahing": 9,
	"Pon":    7,
	"Wage":   4,
	"Kliwon": 8,
}

var pasaran = []string{"Legi", "Pahing", "Pon", "Wage", "Kliwon"}

// Hitung Neptu Weton
func hitungNeptu(hari string, pasaran string) int {
	return neptuHari[hari] + neptuPasaran[pasaran]
}

func coupleMatchResult(neptu1, neptu2 int) string {
	totalNeptu := neptu1 + neptu2

	neptuCategory := map[int]string{
		9:  "Cocok sekali, hubungan harmonis.",
		10: "Kurang harmonis, tetapi bisa diatasi dengan pengertian.",
		12: "Hubungan baik, tetapi perlu perhatian lebih.",
		8:  "Kurang cocok, sering ada salah paham.",
		18: "Kurang cocok, sering ada salah paham.",
		19: "Cocok sekali, hubungan harmonis.",
		20: "Kurang harmonis, tetapi bisa diatasi dengan pengertian.",
		22: "Hubungan baik, tetapi perlu perhatian lebih.",
	}

	if result, exists := neptuCategory[totalNeptu]; exists {
		return result
	}

	// Jika tidak ada, cari neptu terdekat
	nearestNeptu := 0
	closestDiff := int(^uint(0) >> 1) // Nilai besar untuk inisialisasi
	for k := range neptuCategory {
		diff := abs(totalNeptu - k)
		if diff < closestDiff {
			closestDiff = diff
			nearestNeptu = k
		}
	}

	return neptuCategory[nearestNeptu]
}

// Fungsi untuk menghitung selisih absolut (abs)
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
	Hari      string `json:"hari"`
	Pekerjaan string `json:"pekerjaan"`
	Pasaran   string `json:"pasaran"`
}

var careerMap = map[string][]string{
	"Legi": {
		"Guru atau pengajar, karena karakter Legi yang tenang dan bijaksana.",
		"Konsultan atau penasihat, cocok dengan sifatnya yang penyabar dan perhatian.",
		"Pelayan publik atau pegawai administrasi, dengan sikapnya yang sopan dan tertata.",
		"Spiritualis atau pemuka agama, sesuai dengan sifatnya yang religius dan penuh kedamaian.",
		"Terapis atau psikolog, karena kemampuannya dalam mendengarkan dan memberikan solusi.",
	},
	"Pahing": {
		"Pengusaha atau pemimpin, karena sifatnya yang ambisius dan penuh tekad.",
		"Atlet atau pelatih, cocok dengan energi dan semangatnya yang tinggi.",
		"Militer atau polisi, sesuai dengan karakternya yang disiplin dan tangguh.",
		"Motivator atau pembicara publik, berkat kepercayaan diri dan kekuatan bicaranya.",
		"Manajer proyek, karena kemampuannya dalam memimpin dan mengorganisir.",
	},
	"Pon": {
		"Akuntan atau auditor, sesuai dengan sifatnya yang teliti dan detail.",
		"Politikus atau diplomat, cocok dengan kemampuan strategis dan diplomasi Pon.",
		"Pengembang properti atau arsitek, karena kemampuannya dalam merencanakan dengan matang.",
		"Konsultan keuangan, karena kemampuan mengatur dan perhitungan yang cermat.",
		"Analis data atau peneliti, berkat ketajaman berpikirnya.",
	},
	"Wage": {
		"Seniman atau musisi, karena kreativitas tinggi dan sifatnya yang inovatif.",
		"Penulis atau jurnalis, berkat kemampuannya dalam menuangkan ide-ide baru.",
		"Desainer grafis, cocok dengan kemampuan artistiknya dalam menciptakan visual.",
		"Pekerja startup, karena keinginan untuk mencoba hal baru dan berinovasi.",
		"Penemu atau inovator, sesuai dengan sifatnya yang suka berpikir di luar kotak.",
	},
	"Kliwon": {
		"Psikolog atau konselor, berkat kepekaannya terhadap perasaan orang lain.",
		"Penyembuh atau ahli terapi alternatif, cocok dengan intuisi dan empatinya.",
		"Guru spiritual atau mentor, karena kecenderungannya pada pengembangan jiwa.",
		"Perawat atau pekerja sosial, sesuai dengan sifatnya yang penuh perhatian dan welas asih.",
		"Ahli lingkungan atau konservasionis, karena kedekatannya dengan alam dan keseimbangan.",
	},
}

// GetRandomCareer returns a random career suggestion for a given pasaran.
func GetRandomCareer(pasaran string) string {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator
	careers := careerMap[pasaran]

	if len(careers) == 0 {
		return "Tidak ada informasi tersedia."
	}

	return careers[rand.Intn(len(careers))]
}

// SetCORS adds CORS headers to the response
func SetCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")    // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET") // Allow GET method
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// GetWetonHandler handles the /weton/{tanggal-masehi} endpoint
func GetWetonHandler(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	SetCORS(w)
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
			Hari:      gregorianDay,
			Pekerjaan: GetRandomCareer(pasaranDay),
			Pasaran:   pasaranDay,
		},
	}

	// Set response headers and return JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Fungsi untuk mengubah input string menjadi struct time
func parseDate(dateStr string) (time.Time, error) {
	layout := "02-01-2006"
	return time.Parse(layout, dateStr)
}

// Fungsi API untuk mencocokkan weton jodoh
func GetWetonJodohHandler(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	SetCORS(w)
	pria := r.URL.Query().Get("pria")
	wanita := r.URL.Query().Get("wanita")

	if pria == "" || wanita == "" {
		http.Error(w, "Parameter 'pria' dan 'wanita' harus diisi", http.StatusBadRequest)
		return
	}

	tanggalPria, err := parseDate(pria)
	if err != nil {
		http.Error(w, "Format tanggal pria tidak valid. Gunakan dd-mm-yyyy", http.StatusBadRequest)
		return
	}

	tanggalWanita, err := parseDate(wanita)
	if err != nil {
		http.Error(w, "Format tanggal wanita tidak valid. Gunakan dd-mm-yyyy", http.StatusBadRequest)
		return
	}

	// Hitung neptu untuk pria dan wanita
	hariPria := tanggalPria.Weekday().String()
	pasaranPria := GetPasaran(tanggalPria)
	neptuPria := hitungNeptu(hariPria, pasaranPria)

	hariWanita := tanggalWanita.Weekday().String()
	pasaranWanita := GetPasaran(tanggalWanita)
	neptuWanita := hitungNeptu(hariWanita, pasaranWanita)

	// Tentukan kecocokan jodoh
	hasilJodoh := coupleMatchResult(neptuPria, neptuWanita)

	// Response JSON
	response := map[string]interface{}{
		"data": map[string]string{
			"hasil": hasilJodoh,
		},
		"metadata": map[string]string{},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Define the route handler
	http.HandleFunc("/api/weton/", GetWetonHandler)
	http.HandleFunc("/api/jodoh", GetWetonJodohHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on port 7723...")
	http.ListenAndServe(":7723", nil)
}
