package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	c "github.com/eptaccio/willtg/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

// Response on action
type Response events.APIGatewayProxyResponse

// Request action
type Request events.APIGatewayProxyRequest

// Handler http requests
func Handler(ctx context.Context, req Request) (Response, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:       os.Getenv("BOT_TOKEN"),
		Synchronous: true,
	})
	if err != nil {
		panic(err)
	}

	b.Handle(tb.OnText, c.Will)

	var u tb.Update
	if err = json.Unmarshal([]byte(req.Body), &u); err == nil {
		b.ProcessUpdate(u)
	}

	resp := Response{
		StatusCode: 200,
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
