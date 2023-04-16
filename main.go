package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"Be the change that you wish to see in the world. - Mahatma Gandhi",
	"The only thing necessary for the triumph of evil is for good men to do nothing. - Edmund Burke",
	"The best way to predict the future is to create it. - Peter Drucker",
	"Success is not final, failure is not fatal: it is the courage to continue that counts. - Winston Churchill",
	"The only true wisdom is in knowing you know nothing. - Socrates",
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Generate a random index
		index := rand.Intn(len(quotes))
		log.Println("Service called: index = ", index)

		// Get a random quote
		quote := quotes[index]
		log.Println("Quote is:", quote)

		// Write the quote to the response writer
		_, err := fmt.Fprint(w, quote)
		if err != nil {
			log.Printf("Failed to write response: %v", err)
			http.Error(w, "Failed to generate quote", http.StatusInternalServerError)
		}
	})

	// Start the HTTP server
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
