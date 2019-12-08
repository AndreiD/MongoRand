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

func main() {

	db, err := initDatabase(mongoURI, dbName)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	}
	log.Println("Connected ok to the database")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for i := 0; i < runs; i++ {

		name, _ := faker.GetPerson().FirstName(reflect.Value{})
		phone, _ := faker.GetPhoner().PhoneNumber(reflect.Value{})
		website, _ := faker.GetNetworker().URL(reflect.Value{})
		ip, _ := faker.GetNetworker().IPv4(reflect.Value{})
		paragraph, _ := faker.GetLorem().Paragraph(reflect.Value{})

		_, err := db.Collection("users").InsertOne(
			ctx,
			bson.D{
				{"name", name},
				{"phone", phone},
				{"website", website},
				{"ip", ip},
				{"created_at", time.Now().Unix()},
				{"paragraph", paragraph},
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
