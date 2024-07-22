package database

import (
	"context"
	"fmt"
	"os"

	"github.com/BAXMAY/food-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDBConnection() (*mongo.Database, error) {

	mongoDBConnURL, err := utils.ConnectionURLBuilder("mongodb")
	if err != nil {
		return nil, err
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoDBConnURL).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(
		context.TODO(), bson.D{{Key: "ping", Value: 1}},
	).Decode(&result); err != nil {
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()

		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	db := client.Database(os.Getenv("DB_NAME"))

	return db, nil
}
