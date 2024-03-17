package api

import (
    "backend/db"
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// GetGeospatialData retrieves geospatial data from the database
func GetGeospatialData(w http.ResponseWriter, r *http.Request) {
    
    data, err := db.GetGeospatialData()
    if err != nil {
        log.Println("Error retrieving geospatial data:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Respond with the fetched data
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

// UploadGeospatialData handles the upload of geospatial data to the database
func UploadGeospatialData(w http.ResponseWriter, r *http.Request) {
    // Parse the request body
    var data struct {
        Name string `json:"name"`
        Data string `json:"data"`
    }
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Insert data into the database
    err = db.InsertGeospatialData(data.Name, data.Data)
    if err != nil {
        log.Println("Error inserting data into database:", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Respond with success message
    response := struct {
        Message string `json:"message"`
    }{
        Message: "Geospatial data uploaded successfully",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// NewRouter creates a new router with API endpoints
func NewRouter() http.Handler {
    router := mux.NewRouter()

    router.HandleFunc("/api/geospatial", GetGeospatialData).Methods("GET")
    router.HandleFunc("/api/upload", UploadGeospatialData).Methods("POST")

    return router
}
