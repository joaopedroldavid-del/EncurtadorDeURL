package api

import (
	"EncurtadorDeURL/internal/store"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string
}

func handleShortenURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req shortenURLRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			sendJSON(
				w,
				apiResponse{Error: "Invalid request body"},
				http.StatusUnprocessableEntity,
			)
			return
		}

		if _, err := url.Parse(req.URL); err != nil {
			sendJSON(
				w,
				apiResponse{Error: "Invalid URL"},
				http.StatusBadRequest,
			)
			return
		}

		code, err := store.SaveShortenedURL(r.Context(), req.URL)
		if err != nil {
			slog.Error("Failed to create code", "error", err)
			sendJSON(
				w,
				apiResponse{Error: "Internal server error"},
				http.StatusInternalServerError,
			)
			return
		}

		sendJSON(
			w,
			apiResponse{Data: shortenURLResponse{Code: code}},
			http.StatusCreated,
		)
	}
}