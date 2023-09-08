package config

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/bookkyjung1221/pokemon/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Seed() {

	jsonData, err := os.ReadFile("pokemon-init-db.json")
	if err != nil {
		log.Fatal(err)
	}

	var pokemon []models.Pokemon
	if err := json.Unmarshal(jsonData, &pokemon); err != nil {
		log.Fatal(err)
	}

	collection := GetCollection()

	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if count > 0 {
		log.Println("Data already seeded in the MongoDB collection.")
	} else {
		for _, onePokemon := range pokemon {
			_, err := collection.InsertOne(context.Background(), onePokemon)
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Println("Data inserted into MongoDB collection successfully!")
	}
}
