package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RianNegreiros/secure-password/validator"
)

// PasswordRequest is the structure for the password request
type PasswordRequest struct {
	Password string `json:"password"`
}

// ValidatePasswordHandler handles requests to validate passwords
func ValidatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check Content-Type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "The content type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	var req PasswordRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding JSON request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.ValidatePassword(req.Password)
	if err != nil {
		http.Error(w, "Password validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
