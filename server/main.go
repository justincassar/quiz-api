package main

import (
	"QuizAPI/handlers"
	"log"
	"net/http"
)

// Main function
func main() {

	// Set routes for handlers
	http.HandleFunc("/questions/", handlers.HandleQuestion)
	http.HandleFunc("/answer/", handlers.HandleAnswer)
	http.HandleFunc("/result/", handlers.HandleResult)

	//Send all questions data to server (Not implemented)
	//http.HandleFunc("/questions", handleQuestions)

	// Start server. Handler is set to "nil" to use the routers set above
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
