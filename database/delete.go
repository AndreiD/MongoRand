package database
//
//import "go.mongodb.org/mongo-driver/mongo"
//
//func RemoveOneHero(client *mongo.Client, filter bson.M) int64 {
//	collection := client.Database("civilact").Collection("heroes")
//	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
//	if err != nil {
//		log.Fatal("Error on deleting one Hero", err)
//	}
//	return deleteResult.DeletedCount
//}