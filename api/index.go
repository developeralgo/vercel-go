package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	IP      string              `json:"ip"`
	Headers map[string][]string `json:"headers"`
	Method  string              `json:"method"`
	URL     string              `json:"url"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	headers := make(map[string][]string)
	for name, values := range r.Header {
		headers[name] = values
	}

	response := Response{
		IP:      ip,
		Headers: headers,
		Method:  r.Method,
		URL:     r.URL.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
