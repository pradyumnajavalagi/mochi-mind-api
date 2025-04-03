package models

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

type Flashcard struct {
	ID            int    `json:"id"`
	KanjiImageURL string `json:"kanji_image_url"`
	Onyomi        string `json:"onyomi"`
	Kunyomi       string `json:"kunyomi"`
	ExampleUsage  string `json:"example_usage"`
}

var db *pgx.Conn

func InitDB() {
	var err error
	db, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
}

func InsertFlashcard(card Flashcard) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO flashcards (kanji_image_url, onyomi, kunyomi, example_usage)
		VALUES ($1, $2, $3, $4)
	`, card.KanjiImageURL, card.Onyomi, card.Kunyomi, card.ExampleUsage)
	return err
}

func GetAllFlashcards() ([]Flashcard, error) {
	rows, err := db.Query(context.Background(), `SELECT id, kanji_image_url, onyomi, kunyomi, example_usage FROM flashcards`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []Flashcard
	for rows.Next() {
		var c Flashcard
		err = rows.Scan(&c.ID, &c.KanjiImageURL, &c.Onyomi, &c.Kunyomi, &c.ExampleUsage)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}

func GetRandomFlashcards() ([]Flashcard, error) {
	rows, err := db.Query(context.Background(), `SELECT id, kanji_image_url, onyomi, kunyomi, example_usage FROM flashcards ORDER BY RANDOM()`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []Flashcard
	for rows.Next() {
		var c Flashcard
		err = rows.Scan(&c.ID, &c.KanjiImageURL, &c.Onyomi, &c.Kunyomi, &c.ExampleUsage)
		if err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}
	return cards, nil
}

func UpdateFlashcard(card Flashcard) error {
	_, err := db.Exec(context.Background(), `
		UPDATE flashcards SET kanji_image_url=$1, onyomi=$2, kunyomi=$3, example_usage=$4 WHERE id=$5
	`, card.KanjiImageURL, card.Onyomi, card.Kunyomi, card.ExampleUsage, card.ID)
	return err
}

func DeleteFlashcard(id int) error {
	_, err := db.Exec(context.Background(), `DELETE FROM flashcards WHERE id=$1`, id)
	return err
}
