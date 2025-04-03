package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"mochi-mind-api/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	var card models.Flashcard
	json.NewDecoder(r.Body).Decode(&card)
	err := models.InsertFlashcard(card)
	if err != nil {
		http.Error(w, "Unable to insert flashcard", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllFlashcards(w http.ResponseWriter, r *http.Request) {
	cards, err := models.GetAllFlashcards()
	if err != nil {
		http.Error(w, "Unable to get flashcards", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cards)
}

func GetRandomFlashcards(w http.ResponseWriter, r *http.Request) {
	cards, err := models.GetRandomFlashcards()
	if err != nil {
		http.Error(w, "Unable to get flashcards", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cards)
}

func UpdateFlashcard(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var card models.Flashcard
	json.NewDecoder(r.Body).Decode(&card)
	card.ID = id

	err := models.UpdateFlashcard(card)
	if err != nil {
		http.Error(w, "Unable to update flashcard", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteFlashcard(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	err := models.DeleteFlashcard(id)
	if err != nil {
		http.Error(w, "Unable to delete flashcard", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func UploadImage(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Cannot read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), header.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	// Assuming your Flutter app can access this path
	fmt.Fprintf(w, "http://localhost:8080/%s", filePath)
}
