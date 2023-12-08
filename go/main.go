package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response struct to represent JSON response
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Handler function for the API endpoint
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Create a Response struct
		response := Response{Message: "Welcome to the API endpoint!"}

		// Convert the Response struct to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set content type to JSON
		w.Header().Set("Content-Type", "application/json")
		// Write the JSON response
		w.Write(jsonResponse)
	}

	// Define the route and handler for the API endpoint
	http.HandleFunc("/test", handler)

	// Start the HTTP server on port 8080
	fmt.Println("Server running on http://localhost:8080/test")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
