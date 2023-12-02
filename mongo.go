package backprofile

import (
	"context"
    "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database("PakArbi") // menyimpan data di mongoDB
	return db, nil
}

