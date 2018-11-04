package repository

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"gojek-1st/pkg/entity"
	"log"
)

const (
	collectionName = "driver"
)

type MongoRepository struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client, db string) *MongoRepository {
	return &MongoRepository{
		client:     client,
		db:         db,
		collection: client.Database(db).Collection(collectionName),
	}
}

func (mr *MongoRepository) Store(d *entity.Driver) (uint8, error) {

	_, err := mr.collection.InsertOne(context.Background(), d)
	if err != nil {
		log.Fatal(err)
	}
	return d.Id, err
}

func (mr *MongoRepository) Get(id string) (d *entity.Driver, err error) {
	return nil, nil
}
