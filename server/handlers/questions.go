package handlers

import (
	"QuizAPI/data"
	"QuizAPI/utils"
	"QuizAPI/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Handles request for all questions within the database.
// *This was not implemented in the final version and hence commented out.
//func handleQuestions(w http.ResponseWriter, r *http.Request) {
//	encodeData(w, Questions)
//}

// Handles single question requests from the client
func HandleQuestion(w http.ResponseWriter, r *http.Request) {

	// Print starting statement to terminal
	log.Println("Question Request: Client request received")

	// Split URL to obtain ID of question
	id := strings.SplitAfter(r.URL.String(), "/questions/")

	// If the ID is empty, stop operation
	if len(id) < 2 {
		utils.SendBadRequest(w)
		return
	}

	// Error handle int
	intId, err := strconv.Atoi(id[1])
	if err != nil {
		fmt.Println(err.Error())
	}

	// Check if question ID is within range
	if intId < 0 || intId >= len(data.Questions) {

		// Send -1 ID question struct to indicate the ID is out of range
		var stopLoop = models.Question{Id: -1, Question: "", Choices: []string{}}
		utils.EncodeData(w, stopLoop)

	} else {

		// Print output statement
		fmt.Println("Sending question to client - Question ID: ", intId)

		// Search for question using the ID
		fetchQuestion(w, intId)
	}

	// Print success statement
	log.Printf("SUCCESS: Question %v sent to client \n", intId)
}

// Search for question ID and send its data to client
func fetchQuestion(w http.ResponseWriter, requestedId int) {

	// Loop through questions and find related ID
	for _, question := range data.Questions {
		if question.Id == requestedId {

			// Send question to client
			utils.EncodeData(w, question)
			return
		}
	}
}
