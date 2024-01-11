package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/sami-alex/netflix-go-backend/db"
	"github.com/sami-alex/netflix-go-backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	UserName string             `json:"userName" `
	Password string             `json:"password" binding:"required"`
	Email    string             `json:"email" binding:"required"`
}

func (u *User) CreateUser() (any, error) {
	collection := db.Client.Database("netflix").Collection("user")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)

	if err != nil {
		return nil, err
	}
	generatedId := primitive.NewObjectID()

	for {
		err := collection.FindOne(context.TODO(), bson.M{"_id": generatedId})
		if err != nil {
			generatedId = primitive.NewObjectID()

		}
		break
	}
	u.ID = generatedId
	result, err := collection.InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (u *User) Login() error {
	var user User
	collection := *db.Client.Database("netflix").Collection("user")
	err := collection.FindOne(context.TODO(), bson.M{"email": u.Email}).Decode(&user)
	if err != nil {
		fmt.Println(err, "collection")
		return errors.New("Invalid email or password")
	}
	isValid := utils.DecryptPassword(u.Password, user.Password)

	if !isValid {
		return errors.New("Invalid email or password")
	}
	*u = user
	return nil
}
