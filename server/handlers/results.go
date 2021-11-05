package handlers

import (
	"QuizAPI/data"
	"QuizAPI/utils"
	"QuizAPI/models"
	"log"
	"net/http"
)

// Handles results sent from the client
func HandleResult(w http.ResponseWriter, r *http.Request) {

	// Print starting statement to terminal
	log.Println("Result Request: Client request received")

	// Calculate the percentage of users who got a worse result
	worsePercent := calculateWorsePercentage(data.CorrectAnswers)

	// Store user result along with previously obtained results
	data.StoredResults = append(data.StoredResults, data.CorrectAnswers)

	// Create Result slice with number of correct answers and percentage of worse scores
	result := models.Result{Result: data.CorrectAnswers, Percentage: worsePercent}

	// Send data to client
	utils.EncodeData(w, result)

	// Reset results global variables
	resetScores()
}

// Calculate the percentage of worse results
func calculateWorsePercentage(correctAnswers int) int {

	// Initialise variable
	var totalWorse = 0

	// Loop through the quiz`s previous results to check how many got worse results
	for i := 0; i < len(data.StoredResults); i++ {
		if correctAnswers > data.StoredResults[i] {
			totalWorse++
		}
	}

	// Check if user is first to do the quiz and return the 0% immediately. This avoids diving by 0
	if len(data.StoredResults) < 1 {
		return 100
	}

	// Calculate the percentage of worse players
	worsePercent := int((float64(totalWorse) / float64(len(data.StoredResults))) * 100)

	// Return value
	return worsePercent
}

// Reset global variables in preparation for the consecutive game
func resetScores() {
	data.CorrectAnswers = 0
	data.ValidAnswers = 0
}
