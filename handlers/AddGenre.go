package handlers

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/Deep18501/hotstar_mini/models"
)


func (h handler) AddGenre(w http.ResponseWriter, r *http.Request) {
	genre := models.Genre{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&genre)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	result := h.DB.Create(&genre)
	if result.Error != nil {
		log.Printf("Error :%s", result.Error.Error())
		respondWithError(w, 400, "Error Creating genre")
		return
	}
	respondWithJSON(w, 201, genre)
}