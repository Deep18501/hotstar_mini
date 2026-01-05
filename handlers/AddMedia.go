package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/Deep18501/hotstar_mini/models"
)

func (h handler) AddMedia(w http.ResponseWriter, r *http.Request) {
	// Define a temporary struct to map the incoming JSON
	var req struct {
		models.Media
		GenreNames []string `json:"genre"`
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithError(w, 400, "Error reading request body")
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	media := req.Media

	// Process genres
	var genres []*models.Genre
	for _, gName := range req.GenreNames {
		var g models.Genre
		// Find existing genre or create a new one
		if err := h.DB.FirstOrCreate(&g, models.Genre{GenreType: gName}).Error; err != nil {
			log.Printf("Error finding/creating genre %s: %v", gName, err)
			continue
		}
		genres = append(genres, &g)
	}
	media.Genre = genres

	log.Printf("Media to save: %+v", media)
	log.Printf("Genres to link: %+v", req.GenreNames)

	result := h.DB.Create(&media)
	if result.Error != nil {
		log.Printf("Error creating media: %s", result.Error.Error())
		respondWithError(w, 400, "Error Creating media")
		return
	}

	respondWithJSON(w, 201, media)
}
