/*
Copyright © 2021 Justin Cassar

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the quiz",
	Long: `The command will start the displaying the questions with the multiple choices. 
	Select an answer each time`,
	Run: func(cmd *cobra.Command, args []string) {
		startQuiz()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Root URL of the server
var rootURL = "http://localhost:8080"

// Structs used for data handling
type Question struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
}

type Answer struct {
	QuestionId int `json:"questionid"`
	Answer int `json:"answer"`
}

type Result struct {
	Result     int `json:"result"`
	Percentage int `json:"percentage"`
}

// Queries the server for a question using its attributed ID
func getQuestion(id int) Question {

	// Declare variable to store question
	question := Question{}

	// Convert id variable from int to string
	var questionID = strconv.Itoa(id)

	// Concatenate the string for the complete URL
	urlString := rootURL + "/questions/" + questionID

	// Send Get request to server using the URL String
	res, err := http.Get(urlString)
	if err != nil {
		log.Panic(err)
	}

	// Defer the closure of the response body until the function is complete
	defer res.Body.Close()

	// Decode the question sent by the server
	json.NewDecoder(res.Body).Decode(&question)

	// Deliver the question as a result of the function
	return question
}

// Display the question received
func printQuestion(question Question) {

	// Display question text
	fmt.Println("QUESTION: ", question.Question)

	// Display the choices attributed with the question
	if question.Choices != nil {
		for i := range question.Choices {
			fmt.Printf("(%v): %v\n", i+1, question.Choices[i])
		}
	}
}

// Read user's input
func readUserInput(question Question) int {

	// Ask user to type in the answer of choice
	for {
		fmt.Print("Type in number of your chosen answer: ")

		// Start scanner to await user input
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		// Store user input
		userInput := input.Text()

		// Convert user input to integer and validate
		userChoice, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("ERROR: Please input a number within the range of choices")
		} else if userChoice <= 0 || userChoice > len(question.Choices) {
			fmt.Println("ERROR: Type in a number within the range of choices")
		} else {

			// User's choice is returned after validation
			return userChoice-1
			break
		}
	}

	// Return an invalid output in case of function failure
	return -1
}

// Send user's answer to server
func sendAnswer(questionId int, result int) {

	// Store question ID and choice of answer as an instance of Answer struct
	answer := Answer{questionId, result}

	// Convert answer to JSON
	jsonData, _ := json.Marshal(answer)

	// Perform POST request to send JSON data to server
	res, er := http.Post(rootURL+"/answer/", "application/json", bytes.NewBuffer(jsonData))
	if er != nil {
		log.Fatal(er)
	}

	// Defer the closure of the response body until the function is complete
	defer res.Body.Close()

	// Read response body
	_, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
}

// Request the result from the server
func getResults() {

	// Declare empty instance of Result struct
	userResult := Result{}

	// Send Get request to server for the result
	res, err := http.Get(rootURL + "/result/")
	if err != nil {
		log.Panic(err)
	}

	// Defer closure of response body
	defer res.Body.Close()

	// Decode the received result data
	json.NewDecoder(res.Body).Decode(&userResult)

	// Display the result received
	fmt.Printf("You got %v questions correct. You were better than %v%% of all quizzers.", userResult.Result, userResult.Percentage)

}

// Start the quiz
func startQuiz() {

	// Display welcome message
	fmt.Println("Welcome to the quiz!")
	fmt.Println("Answer the question by typing in the associated number of the choices shown below each question.")
	fmt.Println("------------------------")
	fmt.Println("Let's start with the first question: ")

	// Initialise variable for loop index
	i := 0

	// Loop through questions from server
	for {
		question := getQuestion(i)

		// Checks if the server sent end of loop question struct
		if question.Id == -1 {
			getResults()
			break
		}

		// Display question to user
		printQuestion(question)

		// Read user input for the question answer
		userAnswer := readUserInput(question)

		// Send answer to server
		sendAnswer(question.Id, userAnswer)

		// Increment index of for loop
		i++
	}
}
