package internal

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/uptrace/bun"
)

func GetInfo(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		shortURL := vars["shortURL"]

		url := &URL{ShortURL: shortURL}

		err := db.NewSelect().Model(url).Where("short_url = ?", shortURL).Scan(context.Background())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res := struct {
			ShortURL  string `json:"short_url"`
			LongURL   string `json:"long_url"`
			Clicks    int64  `json:"clicks"`
			ExpiresAt string `json:"expires_at"`
		}{
			ShortURL:  url.ShortURL,
			LongURL:   url.LongURL,
			Clicks:    url.Clicks,
			ExpiresAt: url.ExpiresAt.Format("2006-01-02 15:04:05"),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
