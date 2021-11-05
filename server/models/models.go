package models

// This file contains all the structs used
// The tags have also been added at the end to eliminate camelcasing for JSON formatting

type Question struct {
	Id       int      `json:"id"`
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
}

type Answer struct {
	QuestionId int `json:"questionid"`
	Answer     int `json:"answer"`
}

type Result struct {
	Result     int `json:"result"`
	Percentage int `json:"percentage"`
}
