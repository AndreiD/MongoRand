package main

import (
	"context"
	"github.com/bxcodec/faker"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"mongorand/configs"
	"mongorand/database"
	"reflect"
	"time"
)

const runs = 2 // <-------- MODIFY ME

var db *mongo.Database
var configuration *configs.ViperConfiguration

func init() {
	configuration = configs.NewConfiguration()
	configuration.Init()
	log.Println("=======================================")
	log.Printf("Starting Mongo Random Filler. Running %d times", runs)
	log.Println("=======================================")

}
func main() {

	db, err := database.InitDatabase(configuration.Get("database.mongoURI"), configuration.Get("database.dbname"))
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
