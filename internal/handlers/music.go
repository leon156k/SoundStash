package handlers

import (
    "net/http"
    "encoding/json"
    "github.com/yourusername/SoundStash/internal/models"
)

// GetMusic handles retrieving music information
func GetMusic(w http.ResponseWriter, r *http.Request) {
    // Example response
    music := models.Music{
        ID:     1,
        Title:  "Example Song",
        Artist: "Example Artist",
        Album:  "Example Album",
        Genre:  "Example Genre",
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(music)
}