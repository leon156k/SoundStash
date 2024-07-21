package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    // Initialize the router
    r := mux.NewRouter()
    
    // Define routes
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/music", MusicHandler) // Example route for handling music

    // Start the server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal("Error starting server: ", err)
    }
}

// HomeHandler handles the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to SoundStash!"))
}

// MusicHandler handles requests for music data
func MusicHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Music API endpoint"))
}