package handlers

import (
	"net/http"
	"strconv"

	"github.com/Deep18501/hotstar_mini/models"
)

func (h handler) GetMediaById(w http.ResponseWriter, r *http.Request) {
	var media models.Media
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, 400, "Invalid media ID")
		return
	}

	err = h.DB.Preload("Genre").Preload("Rating").First(&media, id).Error
	if err != nil {
		respondWithError(w, 404, "Media not found")
		return
	}

	respondWithJSON(w, 200, media)
}
