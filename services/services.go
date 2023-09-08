package services

import (
	"context"
	"strconv"

	"github.com/bookkyjung1221/pokemon/config"
	"github.com/bookkyjung1221/pokemon/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPokemons() ([]models.Pokemon, error) {

	collection := config.GetCollection()

	var pokemons []models.Pokemon
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var pokemon models.Pokemon
		if err := cursor.Decode(&pokemon); err != nil {
			return nil, err
		}
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func GetPokemon(id string) (*models.Pokemon, error) {
	collection := config.GetCollection()

	numID, _ := strconv.Atoi(id)

	var pokemon models.Pokemon

	filter := bson.M{"id": numID}

	err := collection.FindOne(context.Background(), filter).Decode(&pokemon)

	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func CreatePokemon(req *models.CreatePokemonRequest) error {
	collection := config.GetCollection()

	_, err := collection.InsertOne(context.Background(), &req)

	if err != nil {
		return err
	}
	return nil
}

func UpdatePokemon(id string, req *models.UpdatePokemonRequest) error {

	collection := config.GetCollection()

	numID, _ := strconv.Atoi(id)

	filter := bson.M{"id": numID}
	update := bson.M{"$set": &req}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeletePokemon(id string) error {

	collection := config.GetCollection()

	numID, _ := strconv.Atoi(id)

	filter := bson.M{"id": numID}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func DeletePokemons() error {
	collection := config.GetCollection()
	_, err := collection.DeleteMany(context.Background(), bson.M{})
	return err
}
