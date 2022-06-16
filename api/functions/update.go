package main

import (
	"context"
	"encoding/json"
	"main/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestBody struct {
	Value bool `json:"value"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// get value from request body as boolean
	var requestBody RequestBody
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(err.Error()),
			StatusCode: 400,
		}, nil
	}
	objectId, err := primitive.ObjectIDFromHex(request.PathParameters["id"])
	client := db.ConnectToDb()
	coll := client.Database("redwood-notes").Collection("Note")
	filter := bson.D{{"_id", objectId}}
	update := bson.D{{"$set", bson.D{{"done", requestBody.Value}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(result.ModifiedCount),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
