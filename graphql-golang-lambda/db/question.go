package db

type QuestionType string

const (
	QuestionTypeText       QuestionType = "TEXT"
	QuestionTypeNumeric    QuestionType = "NUMERIC"
	QuestionTypePredefined QuestionType = "PREDEFINED"
)

type Question struct {
	Model
	Title    string
	Type     QuestionType
	Color    string
	Sequence int
}
