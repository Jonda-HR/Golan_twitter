package db

import (
	"context"

	"github.com/Jonda-HR/Goland_twitter/v2/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExist(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	collection := db.Collection("user")

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
