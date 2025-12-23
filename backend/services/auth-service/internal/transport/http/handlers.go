package http

import (
	"encoding/json"
	"net/http"
)

type MagicLinkRequest struct {
	Email string `json:"email"`
}

func MagicLinkRequestHandler(w http.ResponseWriter, r *http.Request) {
	var body MagicLinkRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// TODO: generate token + save to DB
	// TODO: send email or log URL

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Magic link sent (or logged)"))
}
