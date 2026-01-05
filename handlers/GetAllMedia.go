package handlers

import (
	"net/http"

	"github.com/Deep18501/hotstar_mini/models"
)

func (h handler) GetAllMedia(w http.ResponseWriter, r *http.Request) {
	var medias []models.Media
	err := h.DB.Model(&models.Media{}).Preload("Genre").Preload("Rating").Find(&medias).Error

	if err != nil {
		respondWithError(w, 401, "Error getting medias")
	}
	respondWithJSON(w, 201, medias)
}
