package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/leon156k/SoundStash/internal/models"
)

// Example music data
var musicData = []models.Music{
    {ID: 1, Title: "Example Song 1", Artist: "Artist 1", Album: "Album 1", Genre: "Genre 1"},
    {ID: 2, Title: "Example Song 2", Artist: "Artist 2", Album: "Album 2", Genre: "Genre 2"},
}

// GetMusic handles retrieving music information
func GetMusic(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(musicData)
}
