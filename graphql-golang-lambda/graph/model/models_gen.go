// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type NewAnswer struct {
	Content    string `json:"content"`
	QuestionID string `json:"questionId"`
}

type NewQuestion struct {
	Title    string       `json:"title"`
	Type     QuestionType `json:"type"`
	Sequence int          `json:"sequence"`
	Color    string       `json:"color"`
}

type Question struct {
	ID       string       `json:"id"`
	Title    string       `json:"title"`
	Type     QuestionType `json:"type"`
	Color    string       `json:"color"`
	Sequence int          `json:"sequence"`
}

type QuestionType string

const (
	QuestionTypeText       QuestionType = "TEXT"
	QuestionTypeNumeric    QuestionType = "NUMERIC"
	QuestionTypePredefined QuestionType = "PREDEFINED"
)

var AllQuestionType = []QuestionType{
	QuestionTypeText,
	QuestionTypeNumeric,
	QuestionTypePredefined,
}

func (e QuestionType) IsValid() bool {
	switch e {
	case QuestionTypeText, QuestionTypeNumeric, QuestionTypePredefined:
		return true
	}
	return false
}

func (e QuestionType) String() string {
	return string(e)
}

func (e *QuestionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = QuestionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid QuestionType", str)
	}
	return nil
}

func (e QuestionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
