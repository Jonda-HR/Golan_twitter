package db

import (
	"context"

	"github.com/Jonda-HR/Goland_twitter/v2/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertSignUp(user models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	collection := db.Collection("user")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)

	return objId.String(), true, nil
}
