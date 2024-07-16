package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/madushas/go-url-shortner/internal"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := internal.Connect()
	defer db.Close()

	context := context.Background()

	if err := internal.CreateSchema(context, db); err != nil {
		log.Fatal("Error creating schema: %v\n", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", internal.CreateShortURL(context, db)).Methods("POST")
	r.HandleFunc("/{shortURL}", internal.RedirectURL(db)).Methods("GET")

	http.Handle("/", r)
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
