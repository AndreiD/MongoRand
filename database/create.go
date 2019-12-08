package database

//
//import (
//	"context"
//	"errors"
//
//	"golang.org/x/crypto/bcrypt"
//
//	"go.mongodb.org/mongo-driver/mongo"
//
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//)
//
//const userCollection = "user"
//
//var (
//	ErrUserAlreadyExit        = errors.New("user exists")
//	nilUser                   = UserModel{}
//	ErrLoginCredentialInvalid = errors.New("invalid email or password")
//)
//
//// UserModel represents individual User
//type UserModel struct {
//	ID       primitive.ObjectID `bson:"_id"`
//	Email    string             `bson:"email"`
//	Password string             `bson:"password"`
//	Username string             `bson:"username"`
//}
//
//// CreateUser will create an user if the user doesn't exist
//func CreateUser(ctx context.Context, model UserModel) (string, error) {
//	var user UserModel
//
//	err := db.Collection(userCollection).FindOne(ctx,
//		bson.M{"$or": []bson.M{{"username": model.Username},
//			{"email": model.Email}}}).Decode(&user)
//
//	if err != nil && err != mongo.ErrNoDocuments {
//		return primitive.NilObjectID.String(), err
//	}
//	if user != nilUser {
//		return primitive.NilObjectID.String(), ErrUserAlreadyExit
//	}
//	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(model.Password),
//		bcrypt.DefaultCost)
//
//	if err != nil {
//		return primitive.NilObjectID.String(), err
//	}
//
//	model.Password = string(encryptedPass)
//	model.ID = primitive.NewObjectID()
//
//	res, err := db.Collection(userCollection).InsertOne(ctx, model)
//
//	if err != nil {
//		return primitive.NilObjectID.String(), err
//	}
//	id := res.InsertedID.(primitive.ObjectID)
//	return id.String(), err
//}
//
//// CheckValidUser returns true if user credentials are correct
//func CheckValidUser(ctx context.Context, loginID string,
//	password string) (bool, error) {
//	var user UserModel
//	err := db.Collection(userCollection).FindOne(ctx,
//		bson.M{"$or": []bson.M{{"username": loginID},
//			{"email": loginID}}}).Decode(&user)
//
//	if err != nil {
//		return false, err
//	}
//	if user == nilUser {
//		return false, ErrLoginCredentialInvalid
//	}
//	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//
//	if err != nil {
//		return false, nil
//	}
//	return true, nil
//}
