package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	if payload != nil {
		response, _ := json.Marshal(payload)
		_, _ = w.Write(response)
	}
}

type httpError struct {
	err     error
	code    int
	message string
}

func handleError(ctx context.Context, w http.ResponseWriter, err httpError) {
	//TODO logging
	message := err.message
	if message == "" {
		message = err.err.Error()
	}
	writeJSONResponse(w, err.code, apiError{Message: message})
}

type apiError struct {
	Message string `json:"error,omitempty"`
}
