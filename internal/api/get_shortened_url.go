package api

import (
	"EncurtadorDeURL/internal/store"
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type GetShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}
	
func handleGetShortenedURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, err := store.GetFullURL(r.Context(), code)
		
		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJSON(
					w,
					apiResponse{Error: "URL not found"},
					http.StatusNotFound,
				)
				return
			}
			slog.Error("Failed to get code", "error", err)
			sendJSON(
				w,
				apiResponse{Error: "Internal server error"},
				http.StatusInternalServerError,
			)
			return
		}
		
		sendJSON(
			w,
			apiResponse{Data: GetShortenedURLResponse{FullURL: fullURL}},
			http.StatusOK,
		)
	}
}