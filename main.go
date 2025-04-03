package main

import (
	"log"
	"net/http"
	"os"

	"mochi-mind-api/models"
	"mochi-mind-api/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	models.InitDB()

	r := router.SetupRouter()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
