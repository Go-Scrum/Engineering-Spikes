package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"goscrum/Engineering-Spikes/graphql-golang-lambda/db"
	"goscrum/Engineering-Spikes/graphql-golang-lambda/graph/generated"
	"goscrum/Engineering-Spikes/graphql-golang-lambda/graph/model"

	"github.com/jinzhu/gorm"
)

func (r *answerResolver) Question(ctx context.Context, obj *model.Answer) (*model.Question, error) {
	//dbClient := db.DbClient(true)
	//defer dbClient.Close()
	var question db.Question
	err := r.DB.Where("id =?", obj.QuestionID).First(&question).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return convertToQuestion(question), nil
}

func convertToQuestion(question db.Question) *model.Question {
	return &model.Question{
		Color:    question.Color,
		Title:    question.Title,
		Sequence: question.Sequence,
		ID:       question.ID,
		Type:     model.QuestionType(question.Type),
	}
}

func convertToAnswer(answer db.Answer) *model.Answer {
	return &model.Answer{
		Content:    answer.Content,
		ID:         answer.ID,
		QuestionID: answer.QuestionID,
	}
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	//dbClient := db.DbClient(true)
	//defer dbClient.Close()
	question := db.Question{
		Title:    input.Title,
		Type:     db.QuestionType(input.Type),
		Color:    input.Color,
		Sequence: input.Sequence,
	}
	err := r.DB.Save(&question).Error
	if err != nil {
		return nil, err
	}
	return convertToQuestion(question), nil
}

func (r *mutationResolver) CreateAnswer(ctx context.Context, input model.NewAnswer) (*model.Answer, error) {
	answer := db.Answer{
		Content:    input.Content,
		QuestionID: input.QuestionID,
	}
	dbClient := db.DbClient(true)
	defer dbClient.Close()
	err := dbClient.Save(&answer).Error
	if err != nil {
		return nil, err
	}
	return convertToAnswer(answer), nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	var questions []db.Question
	//dbClient := db.DbClient(true)
	//defer dbClient.Close()
	err := r.DB.Find(&questions).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	var qlQuestions []*model.Question
	for _, question := range questions {
		qlQuestions = append(qlQuestions, convertToQuestion(question))
	}
	return qlQuestions, nil
}

func (r *queryResolver) Answers(ctx context.Context) ([]*model.Answer, error) {
	var answers []db.Answer
	//dbClient := db.DbClient(true)
	//defer dbClient.Close()
	err := r.DB.Find(&answers).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	var qlAnswers []*model.Answer
	for _, answer := range answers {
		qlAnswers = append(qlAnswers, convertToAnswer(answer))
	}
	return qlAnswers, nil
}

// Answer returns generated.AnswerResolver implementation.
func (r *Resolver) Answer() generated.AnswerResolver { return &answerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type answerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
