package services

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDbClient() (mongoClient *mongo.Client, err error) {

	mongoURI := os.Getenv("MONGODB_URL")

	clientOptions := options.Client().ApplyURI(mongoURI)

	direct := true
	clientOptions.Direct = &direct

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB:", err)
		return client, err
	} else {
		fmt.Print("Succesfully connected to MongoDb!")
	}

	return client, nil

}
