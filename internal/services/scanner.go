package services

import (
    "os"
    "path/filepath"
    "github.com/leon156k/SoundStash/internal/models"
)

// Scanner provides methods to scan directories and extract metadata
type Scanner struct {
    // Add fields if needed
}

// ScanDirectory scans a directory for music files
func (s *Scanner) ScanDirectory(path string) ([]models.Music, error) {
    var musicFiles []models.Music

    err := filepath.Walk(path, func(filePath string, info os.DirEntry, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && (filepath.Ext(filePath) == ".mp3" || filepath.Ext(filePath) == ".flac") {
            // Example extraction logic; replace with actual metadata extraction
            musicFiles = append(musicFiles, models.Music{
                Title:  info.Name(),
                // Populate other fields as needed
            })
        }
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    return musicFiles, nil
}
