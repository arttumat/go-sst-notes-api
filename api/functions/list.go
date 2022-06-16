package main

import (
	"encoding/json"
	"fmt"
	"main/db"

	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Note struct {
	ObjectID primitive.ObjectID `bson:"_id" json:"_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
	Done     bool               `bson:"done"`
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	client := db.ConnectToDb()
	coll := client.Database("redwood-notes").Collection("Note")
	filter := bson.M{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No documents found")
		return events.APIGatewayProxyResponse{
			Body:       `{"error":true}`,
			StatusCode: 404,
		}, nil
	}
	var response []Note
	for _, result := range results {
		// log result
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			panic(err)
		}
		var note Note
		_ = json.Unmarshal([]byte(output), &note)
		response = append(response, note)
	}
	jsonData, err := json.MarshalIndent(response, "", "    ")

	return events.APIGatewayProxyResponse{
		Body:       string(jsonData),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
