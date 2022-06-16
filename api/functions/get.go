package main

import (
	"context"
	"encoding/json"
	"main/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	client := db.ConnectToDb()
	coll := client.Database("redwood-notes").Collection("Note")
	objectId, _ := primitive.ObjectIDFromHex(request.PathParameters["id"])
	filter := bson.D{{"_id", objectId}}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return events.APIGatewayProxyResponse{
				Body:       `{"error":true}`,
				StatusCode: 404,
			}, nil
		}
		panic(err)
	}

	output, err := json.MarshalIndent(result, "", "    ")

	return events.APIGatewayProxyResponse{
		Body:       string(output),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
