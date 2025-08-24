package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	IP      string              `json:"ip"`
	Headers map[string][]string `json:"headers"`
	Method  string              `json:"method"`
	URL     string              `json:"url"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	// write the whole request including the headers to the response and return as json object
	ip := r.RemoteAddr

	// Create a map to hold the headers
	headers := make(map[string][]string)
	for name, values := range r.Header {
		headers[name] = values
	}

	// Create the response object
	response := Response{
		IP:      ip,
		Headers: headers,
		Method:  r.Method,
		URL:     r.URL.String(),
	}

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response as JSON and send it
	json.NewEncoder(w).Encode(response)

}
