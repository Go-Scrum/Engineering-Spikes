package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack/slackevents"
)

var ginLambda *ginadapter.GinLambda

func init() {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()

	r.POST("/slack/event", func(c *gin.Context) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request.Body)
		body := buf.String()
		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: "dt91zlZzuqgmqErlhPZqvnat"}))
		//spew.Dump(eventsAPIEvent)
		if e != nil {
			fmt.Println(e.Error())
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		if eventsAPIEvent.Type == slackevents.URLVerification {
			fmt.Println("Verification URL sent")
			var r *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)
			if err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{})
			}
			c.Data(200, "text", []byte(r.Challenge))
			return
		}

		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *slackevents.MessageEvent:
				spew.Dump(ev)
				//api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
				//c.JSON(http.StatusOK, gin.H{})
			}
		}

		//////eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(c.Request.GetBody()), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: "TOKEN"}))
		////if e != nil {
		////	w.WriteHeader(http.StatusInternalServerError)
		////}
		//spew.Dump(eventsAPIEvent)
		c.JSON(http.StatusOK, gin.H{})
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

//Handler for handling slack comments
//func Handler() {
//	slackClient := slack.New(os.Getenv("BOT_TOKEN"))
//	start := time.Now()
//	users, _ := slackClient.GetUsers()
//	var wg sync.WaitGroup
//	wg.Add(len(users))
//	for _, user := range users {
//		go func(usr slack.User) {
//			if !usr.IsBot {
//				fmt.Printf("ID: %s TZ: %s Name: %s\n", usr.ID, usr.TZ, usr.Name)
//				_, _, channelID, err := slackClient.OpenIMChannel(usr.ID)
//				if err != nil {
//					fmt.Printf("%s\n", err)
//				}
//				_, _, err = slackClient.PostMessage(channelID, slack.MsgOptionText("Hello world", false))
//				if err != nil {
//					fmt.Printf("%s\n", err)
//				}
//			}
//			defer wg.Done()
//		}(user)
//	}
//	wg.Wait()
//	fmt.Println("Finished for loop")
//	elapsed := time.Since(start)
//	log.Printf("Binomial took %s", elapsed)
//}
