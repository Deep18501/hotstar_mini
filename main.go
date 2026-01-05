package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Deep18501/hotstar_mini/db"
	"github.com/Deep18501/hotstar_mini/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in environemnt")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("dbURL not found in environemnt")
	}

	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/media", h.GetAllMedia).Methods(http.MethodGet)
	router.HandleFunc("/media", h.AddMedia).Methods(http.MethodPost)
	router.HandleFunc("/upload_media", h.UploadMedia).Methods(http.MethodPost)
	router.HandleFunc("/media_detail", h.GetMediaById).Methods(http.MethodGet)
	router.HandleFunc("/category", h.CreateCategory).Methods(http.MethodPost)
	router.HandleFunc("/category", h.GetAllCategory).Methods(http.MethodGet)
	router.HandleFunc("/genre", h.AddGenre).Methods(http.MethodPost)
	router.HandleFunc("/rating", h.SendRating).Methods(http.MethodPost)

	// Serve static files from the uploads directory
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	log.Println("API is running!")
	err := http.ListenAndServe(":"+portString, router)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
	log.Println("on Port :", portString)

}
