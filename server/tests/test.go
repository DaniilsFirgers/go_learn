package tests

import (
	"context"
	"fmt"
	"go-learn/internal/services"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertDummyDocumentIntoDb() {
	client, err := services.GetMongoDbClient()

	if err != nil {
		fmt.Print("error while initializing mongodb client. ")
		return
	}
	collection := client.Database("go_learn").Collection("dummy")

	// Create a sample document
	document := bson.M{
		"name":  "John Doe",
		"email": "john.doe@example.com",
	}

	// Insert the document
	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		fmt.Println("Failed to insert document:", err)
		return
	}

	fmt.Println("Dummy document inserted successfully")
}
