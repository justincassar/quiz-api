package data

import "QuizAPI/models"

// File contains all stored data to act as a database. Data is formatted in JSON.
// Field names are added before the inputs to assist with maintenance
var Questions = []models.Question{
	{Id: 0, Question: "What does “HTTP” stand for?", Choices: []string{"Hyper Text Protocol", "HyperText Transfer Protocol", "Hyper Text Processes"}},
	{Id: 1, Question: "Where is fuel stored on most commercial aircraft?", Choices: []string{"Landing Gears", "Flight deck", "Vertical Stabiliser", "Wings"}},
	{Id: 2, Question: "Uruguay won the first FIFA World Cup in 1930", Choices: []string{"True", "False"}},
}

var Answers = []models.Answer{
	{QuestionId: 0, Answer: 1},
	{QuestionId: 1, Answer: 3},
	{QuestionId: 2, Answer: 0},
}
