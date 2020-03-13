package main

import (
	"context"
	"fmt"

	"goscrum/Engineering-Spikes/graphql-golang-lambda/db"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(_ context.Context) {
	dbClient := db.DbClient(true)
	defer dbClient.Close()

	err := dbClient.AutoMigrate(
		&db.Question{},
		&db.Answer{}).Error

	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Request Initiated")
	lambda.Start(HandleRequest)
	//HandleRequest(context.Background())
}
