package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "/app/data/numbers.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS numbers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		value REAL NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/numbers", numbersHandler)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func numbersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePostNumber(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostNumber(w http.ResponseWriter, r *http.Request) {
	numberStr := r.FormValue("number")
	if numberStr == "" {
		http.Error(w, "Missing number parameter", http.StatusBadRequest)
		return
	}

	number, err := strconv.ParseFloat(numberStr, 64)
	if err != nil {
		http.Error(w, "Invalid number format", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO numbers (value) VALUES (?)", number)
	if err != nil {
		http.Error(w, "Failed to store number", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT value FROM numbers ORDER BY value")
	if err != nil {
		http.Error(w, "Failed to retrieve numbers", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var numbers []float64
	for rows.Next() {
		var value float64
		if err := rows.Scan(&value); err != nil {
			http.Error(w, "Failed to scan number", http.StatusInternalServerError)
			return
		}
		numbers = append(numbers, value)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}
