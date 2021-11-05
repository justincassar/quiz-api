package handlers

import (
	"QuizAPI/data"
	"QuizAPI/models"
	"encoding/json"
	"log"
	"net/http"
)

// Handles answers received from client
func HandleAnswer(w http.ResponseWriter, r *http.Request) {

	// Print statement to terminal
	log.Println("Answer: Client data received")

	// Store the client data as an Answer struct
	var answerData models.Answer

	// Decode the received JSON data in POST request from client
	err := json.NewDecoder(r.Body).Decode(&answerData)
	if err != nil {
		panic(err)
		log.Println(err.Error())
	}

	// Validate the question id before checking if the answer is correct
	if validateAnswer(answerData) {

		// Check if the answer is correct
		checkCorrect(answerData)

	} else {

		// Display validation error
		log.Println("ERROR: Answer Validation - Question ID is not valid")
	}
}

// Validates if the Answer sent by the client has a question identifier that is attributed to a stored question
func validateAnswer(answer models.Answer) bool {

	// Check for an associated question id within the Question slice
	for _, question := range data.Questions {
		if question.Id == answer.QuestionId {

			// if valid, increase the amount of valid answers given by the player
			data.ValidAnswers++
			return true
		}
	}

	// Return false is the question id does not exist
	return false
}

// Checks if the answer sent by the client is correct and if so, it increments the correct answer counter
func checkCorrect(answer models.Answer) {

	// Loops through Answers slice stored on server
	for _, storedAnswer := range data.Answers {

		// Check if the answer with the associated question id are correct
		if storedAnswer.QuestionId == answer.QuestionId && storedAnswer.Answer == answer.Answer {

			// If the answer is correct, increase correct answer to the player's result
			data.CorrectAnswers++
		}
	}
}
