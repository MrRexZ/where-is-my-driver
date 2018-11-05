package repository

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/replaceopt"
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

func (mr *MongoRepository) Store(d *entity.Driver) (int32, error) {
	_, err := mr.collection.ReplaceOne(nil, bson.NewDocument(
		bson.EC.Int32("id", d.Id),
	), d, replaceopt.Upsert(true))

	if err != nil {
		log.Fatal(err)
	}
	return d.Id, err
}

func (mr *MongoRepository) Get(id string) (d *entity.Driver, err error) {
	return nil, nil
}
