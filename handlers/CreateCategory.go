package handlers

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"

	"github.com/Deep18501/hotstar_mini/models"
)


func (h handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&category)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}
	result := h.DB.Create(&category)
	if result.Error != nil {
		log.Printf("Error :%s", result.Error.Error())
		respondWithError(w, 400, "Error Creating category")
		return
	}
	respondWithJSON(w, 201, category)
}