package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func sendJSON(w http.ResponseWriter, resp apiResponse, status int){
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil{
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			apiResponse{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil{
		slog.Error("Failed to write response to client", "error", err)
		return
	}
}