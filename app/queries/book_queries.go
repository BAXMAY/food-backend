package queries

import "go.mongodb.org/mongo-driver/mongo"

type BookQueries struct {
	*mongo.Database
}
