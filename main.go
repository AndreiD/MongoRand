package main

import (
	"context"
	"github.com/bxcodec/faker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"math/rand"
	"reflect"
	"time"
)

// MODIFY ME
const (
	runs     = 2
	mongoURI = "mongodb+srv://user:pass@xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/test?retryWrites=true&w=majority"
	dbName   = "DB_NAME_HERE"
)

var db *mongo.Database

type RandStruct struct {
	Latitude      float32 `faker:"lat"`
	Longitude     float32 `faker:"long"`
	Email         string  `faker:"email"`
	PhoneNumber   string  `faker:"phone_number"`
	URL           string  `faker:"url"`
	LastName      string  `faker:"last_name"`
	FirstName     string  `faker:"first_name"`
	Paragraph     string  `faker:"paragraph"`
	UUIDHypenated string  `faker:"uuid_hyphenated"`
}

func main() {

	db, err := initDatabase(mongoURI, dbName)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	log.Println("Connected ok to the database")

	for i := 0; i < runs; i++ {

		r := RandStruct{}
		err := faker.FakeData(&r)
		if err != nil {
			log.Fatal(err)
		}

		_, err := db.Collection("users").InsertOne(
			context.Background(),
			bson.D{
				{"name", r.FirstName},
				{"phone", r.PhoneNumber},
				{"website", r.URL},
				{"created_at", time.Now().Unix()},
				{"paragraph", r.Paragraph},
				{"qty", rand.Intn(100)},
				{"expertise", bson.A{"computers", "cars", "planes", "buses"}},
				{"office_hours", bson.D{
					{"Monday", "9:00am-5:00pm"},
					{"Tuesday", "8:00am-6:00pm"},
				}},
			})

		if err != nil {
			log.Panic(err)
		}
		if i%100 == 0 {
			log.Printf("finished %d inserts", i)
		}
	}

}

func initDatabase(mongoURI string, Dbname string) (*mongo.Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client.Database(Dbname), nil
}
