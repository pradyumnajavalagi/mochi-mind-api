package router

import (
	"net/http"
	"mochi-mind-api/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/flashcards", middleware.CreateFlashcard).Methods("POST","OPTIONS")
	r.HandleFunc("/flashcards", middleware.GetAllFlashcards).Methods("GET","OPTIONS")
	r.HandleFunc("/flashcards/random", middleware.GetRandomFlashcards).Methods("GET","OPTIONS")
	r.HandleFunc("/flashcards/{id}", middleware.UpdateFlashcard).Methods("PUT","OPTIONS")
	r.HandleFunc("/flashcards/{id}", middleware.DeleteFlashcard).Methods("DELETE","OPTIONS")
	r.HandleFunc("/upload", middleware.UploadImage).Methods("POST", "OPTIONS")
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))


	return r
}
