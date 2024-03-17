package main

import (
    "backend/api"
    "backend/db"
    "log"
    "net/http"
)

func main() {
    // Initialize database
    db.InitDB()

    // Set up routes
    router := api.NewRouter()

    // Start server
    log.Fatal(http.ListenAndServe(":8080", router))
}
