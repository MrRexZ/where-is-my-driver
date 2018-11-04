package repository

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

func CreateMongoClient(mongoUrl string) (*mongo.Client, error) {
	client, err := mongo.NewClient(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}

func CreateMongoRepository(client *mongo.Client, dbName string) *MongoRepository {
	mongoRepository := NewMongoRepository(client, dbName)
	return mongoRepository
}

func DropDatabase(client *mongo.Client, dbName string) error {
	if client != nil {
		err := client.Database(dbName).Drop(context.Background())
		return err
	}
	return nil
}
