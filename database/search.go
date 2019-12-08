package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"iqlaw/models"
	"iqlaw/utils/log"
)

func Search(location string, expertise string, intOffset int, intLimit int) ([]models.User, error) {
	var users []models.User
	collection := db.Collection("users")

	err := validateSearchInput(location, expertise, intOffset, intLimit)
	if err != nil {
		return users, fmt.Errorf(err.Error())
	}

	filter := bson.D{{"$and", []bson.D{
		{{"fee_currency", "USD"}},
		{{"website", "bulgaria"}}},
	}}

	findOptions := options.Find()
	findOptions.SetLimit(1000)

	cur, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return users, err
	}
	for cur.Next(context.Background()) {
		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			log.Error(err)
			continue
		}
		users = append(users, user)
	}
	err = cur.Close(context.Background())
	if err != nil {
		log.Error(err)
	}

	return users, nil
}

// FindOne ..
func FindOne(filter bson.M) models.User {
	var user models.User
	collection := db.Collection("users")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&user)
	return user
}

func validateSearchInput(location string, expertise string, offset int, limit int) error {
	if location == "" {
		return fmt.Errorf("empty location")
	}
	if expertise == "" {
		return fmt.Errorf("empty expertise")
	}
	if offset < 0 {
		return fmt.Errorf("invalid offset")
	}
	if limit > 60000 {
		return fmt.Errorf("there's a limit on how many records you can query")
	}
	return nil
}
