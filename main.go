package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusInternalServerError)
		return
	}

	// Log the request body to STDOUT
	log.Println(string(body))

	// Send a response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request received and logged successfully"))
}

func main() {
	http.HandleFunc("/", handlePostRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
