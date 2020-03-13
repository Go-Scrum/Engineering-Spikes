package db

type Answer struct {
	Model
	ParticipantID string `db:"participant_id" json:"participant_id"`
	QuestionID    string `db:"question_id" json:"question_id"`
	Content       string `db:"comment" json:"comment"`
	Question      Question
}
