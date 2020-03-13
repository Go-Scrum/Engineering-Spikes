package main

import (
	"context"
	"goscrum/Engineering-Spikes/graphql-golang-lambda/db"
	"time"

	"goscrum/Engineering-Spikes/graphql-golang-lambda/graph"
	"goscrum/Engineering-Spikes/graphql-golang-lambda/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var ginLambda *ginadapter.GinLambda
var dbClient *gorm.DB

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	if dbClient == nil {
		dbClient = db.DbClient(true)
	}
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: dbClient}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	r := gin.Default()
	r.POST("/graphql", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": time.Now(),
		})
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
