package database
//
//import "go.mongodb.org/mongo-driver/mongo"
//
//func UpdateHero(client *mongo.Client, updatedData bson.M, filter bson.M) int64 {
//	collection := client.Database("civilact").Collection("heroes")
//	atualizacao := bson.D{ {Key: "$set", Value: updatedData} }
//	updatedResult, err := collection.UpdateOne(context.TODO(), filter, atualizacao)
//	if err != nil {
//		log.Fatal("Error on updating one Hero", err)
//	}
//	return updatedResult.ModifiedCount
//}