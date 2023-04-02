package repositories

import (
	"context"
	"log"

	"github.com/diegoparra/check-birthday/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, t *models.User, coll *mongo.Collection) (string, error) {
	res, err := coll.InsertOne(ctx, t)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func GetUser(ctx context.Context, coll *mongo.Collection, today string) []models.User {
	var result []models.User
	query := bson.M{"birthday": today}
	cursor, err := coll.Find(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}

	return result
}
