package database

import (
	"os"

	"github.com/BAXMAY/food-backend/app/queries"
	"go.mongodb.org/mongo-driver/mongo"
)

type Queries struct {
	*queries.UserQueries
	*queries.BookQueries
}

func OpenDBConnection() (*Queries, error) {
	var (
		db  *mongo.Database
		err error
	)

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "mongodb":
		db, err = MongoDBConnection()
	}

	if err != nil {
		return nil, err
	}

	return &Queries{
		UserQueries: &queries.UserQueries{Database: db},
		BookQueries: &queries.BookQueries{Database: db},
	}, nil
}
