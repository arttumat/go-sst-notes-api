package main

import (
	"context"
	"encoding/json"
	"main/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
)

type InsertRequestBody struct {
	Title string `json:"title"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody InsertRequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	client := db.ConnectToDb()
	coll := client.Database("redwood-notes").Collection("Note")

	doc := bson.D{{"title", requestBody.Title}, {"done", false}}

	_, err = coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
		return events.APIGatewayProxyResponse{
			Body:       string(err.Error()),
			StatusCode: 400,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
