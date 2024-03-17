package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	// Connect to SQLite database
	database, err := sql.Open("sqlite3", "./geospatial.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// Create tables if they don't exist
	_, err = database.Exec(`
        CREATE TABLE IF NOT EXISTS geospatial_data (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT,
            data TEXT
        );
    `)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	db = database
}

// InsertGeospatialData inserts geospatial data into the database
func InsertGeospatialData(name, data string) error {
	_, err := db.Exec("INSERT INTO geospatial_data (name, data) VALUES (?, ?)", name, data)
	if err != nil {
		log.Println("Error inserting data into database:", err)
		return err
	}
	return nil
}

// GetGeospatialData retrieves geospatial data from the database
func GetGeospatialData() ([]GeospatialData, error) {
	rows, err := db.Query("SELECT id, name, data FROM geospatial_data")
	if err != nil {
		log.Println("Error retrieving geospatial data:", err)
		return nil, err
	}
	defer rows.Close()

	var geospatialData []GeospatialData
	for rows.Next() {
		var data GeospatialData
		err := rows.Scan(&data.ID, &data.Name, &data.Data)
		if err != nil {
			log.Println("Error scanning geospatial data:", err)
			continue
		}
		geospatialData = append(geospatialData, data)
	}

	return geospatialData, nil
}

// GeospatialData represents a geospatial data entry in the database
type GeospatialData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data string `json:"data"`
}
