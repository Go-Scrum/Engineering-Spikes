package model

type Answer struct {
	ID         string `json:"id"`
	Content    string `json:"content"`
	QuestionID string `json:"question"`
}
