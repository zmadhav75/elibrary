package utils

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func ValidateISBN(isbn string) bool {
	// Basic ISBN validation
	matched, _ := regexp.MatchString(`^(?:\d{9}[\dXx]|\d{13})$`, isbn)
	return matched
}
