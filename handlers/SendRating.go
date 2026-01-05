package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/Deep18501/hotstar_mini/models"
)

func (h handler) SendRating(w http.ResponseWriter, r *http.Request) {
	// Define a temporary struct to map the incoming JSON

	rating := models.Rating{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&rating)
		if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}

	result := h.DB.Create(&rating)
	if result.Error != nil {
		log.Printf("Error creating media: %s", result.Error.Error())
		respondWithError(w, 400, "Error Creating media")
		return
	}

	respondWithJSON(w, 201, rating)
}
