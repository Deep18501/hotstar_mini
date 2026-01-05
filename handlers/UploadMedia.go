package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Deep18501/hotstar_mini/models"
)

func (h handler) UploadMedia(w http.ResponseWriter, r *http.Request) {
	// 1. Parse Multipart Form (32MB max memory)
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing multipart form")
		return
	}

	saveFile := func(formKey string) (string, error) {
		file, header, err := r.FormFile(formKey)
		if err != nil {
			if err == http.ErrMissingFile {
				return "", nil // Optional file
			}
			return "", err
		}
		defer file.Close()

		// Create unique filename
		ext := filepath.Ext(header.Filename)
		filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), formKey, ext)
		dstPath := filepath.Join("uploads", filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return "", err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			return "", err
		}
		return dstPath, nil
	}

	thumbPath, err := saveFile("thumbnail")
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error saving thumbnail")
		return
	}
	bannerPath, err := saveFile("banner")
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error saving banner")
		return
	}
	mediaPath, err := saveFile("media")
	if err != nil {
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Error saving media")
		return
	}

	
	media := models.Media{
		Title:        r.FormValue("title"),
		Desc:         r.FormValue("desc"),
		ThumbnailURL: thumbPath,
		BannerURL:    bannerPath,
       	MediaUrl: mediaPath,
		AgeRating:    r.FormValue("age_rating"),
		ReleaseYear:  r.FormValue("release_year"),
	}

	// Handle Category ID
	if catIDStr := r.FormValue("category_id"); catIDStr != "" {
		catID, _ := strconv.Atoi(catIDStr)
		media.CategoryId = &catID
	}

	genreInput := r.FormValue("genre")
	var genres []*models.Genre
	if genreInput != "" {
		genreNames := strings.Split(genreInput, ",")
		for _, name := range genreNames {
			name = strings.TrimSpace(name)
			if name == "" {
				continue
			}
			var g models.Genre
			if err := h.DB.FirstOrCreate(&g, models.Genre{GenreType: name}).Error; err != nil {
				log.Printf("Error processing genre %s: %v", name, err)
				continue
			}
			genres = append(genres, &g)
		}
	}
	media.Genre = genres

	// 6. Save to DB
	if result := h.DB.Create(&media); result.Error != nil {
		log.Printf("Error creating media: %v", result.Error)
		respondWithError(w, http.StatusInternalServerError, "Error saving media to database")
		return
	}

	respondWithJSON(w, http.StatusCreated, media)
}