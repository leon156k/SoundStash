package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/leon156k/SoundStash/internal/handlers"
    "github.com/leon156k/SoundStash/internal/services"
)

func main() {
    r := mux.NewRouter()

    // Serve static files
    r.Handle("/", http.FileServer(http.Dir(".")))

    // Define API routes
    r.HandleFunc("/api/music", handlers.GetMusic)
    r.HandleFunc("/api/scan", ScanHandler)  // New endpoint for scanning

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}

// HomeHandler serves a simple homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to SoundStash!"))
}

// ScanHandler handles directory scanning and returns music files
func ScanHandler(w http.ResponseWriter, r *http.Request) {
    scanner := services.Scanner{}
    path := "." // Adjust as needed

    musicFiles, err := scanner.ScanDirectory(path)
    if err != nil {
        http.Error(w, "Error scanning directory", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(musicFiles)
}
