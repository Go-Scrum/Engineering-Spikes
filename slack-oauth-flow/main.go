package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()

	r.GET("/oauth/callback", func(c *gin.Context) {
		redirectURL := "http://localhost:3000/dev/oauth/callback"
		code := c.Query("code")
		clientID := os.Getenv("CLIENT_ID")
		fmt.Println(clientID)
		clientSecret := os.Getenv("CLIENT_SECRET")
		println(clientSecret)
		authResponse, err := slack.GetOAuthV2Response(&http.Client{}, clientID, clientSecret, code, redirectURL)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, authResponse)
	})

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
