package handlers

import (
	"net/http"

	"github.com/Deep18501/hotstar_mini/models"
)

func (h handler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	err := h.DB.Model(&models.Category{}).Preload("Media").Find(&categories).Error

	if err != nil {
		respondWithError(w, 401, "Error getting categories")
	}
	respondWithJSON(w, 200, categories)
}
