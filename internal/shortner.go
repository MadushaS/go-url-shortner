package internal

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
)

type ShortenRequest struct {
	URL       string    `json:"url"`
	CustomURL string    `json:"custom_url"`
	ExpiresAt time.Time `json:"expires"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreateShortURL(ctx context.Context, db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ShortenRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		shortURL := req.CustomURL

		if shortURL == "" {
			shortURL = generateShortURL()
		}

		url := &URL{
			ShortURL:  shortURL,
			LongURL:   req.URL,
			ExpiresAt: req.ExpiresAt,
		}

		_, err = db.NewInsert().Model(url).Exec(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := ShortenResponse{
			ShortURL: os.Getenv("BASE_URL") + shortURL,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func RedirectURL(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		url := &URL{ShortURL: shortURL}
		err := db.NewSelect().Model(url).Where("short_url = ?", shortURL).Scan(context.Background())

		if err != nil {
			http.NotFound(w, r)
			return
		}

		if !url.ExpiresAt.IsZero() && url.ExpiresAt.Before(time.Now()) {
			http.Error(w, "URL has expired", http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, url.LongURL, http.StatusSeeOther)
	}
}

func generateShortURL() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)

	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortURL)
}
