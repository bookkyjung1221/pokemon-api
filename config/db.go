package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() error {

	uri := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(uri)

	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	DB = mongoClient.Database("pokedex")

	log.Println("Connecting to database")

	return nil
}

func GetCollection() *mongo.Collection {
	return DB.Collection("pokemon")
}

func Close() error {
	return DB.Client().Disconnect(context.Background())
}
