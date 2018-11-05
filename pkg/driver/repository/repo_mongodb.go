package repository

import (
	"context"
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

func (mr *MongoRepository) StoreMany(ds []*entity.Driver) error {
	var err error
	for _, v := range ds {
		_, err = mr.Store(v)
	}
	return err
}

func (mr *MongoRepository) Get(id int32) (d *entity.Driver, err error) {
	driver := entity.Driver{}
	doc_res := mr.collection.FindOne(nil,
		bson.NewDocument(bson.EC.Int32("id", id)))
	doc_res.Decode(&driver)
	return &driver, err
}

func (mr *MongoRepository) GetAll() (ds []*entity.Driver, err error) {
	cursor, err := mr.collection.Find(nil, nil)
	var drivers []*entity.Driver
	for cursor.Next(context.Background()) {
		driver := entity.Driver{}
		err = cursor.Decode(&driver)
		drivers = append(drivers, &driver)
	}
	return drivers, err
}

func (mr *MongoRepository) GetWithinLatLng(lat float64, long float64, dist float64) ([]*entity.Driver, error) {
	return nil, nil
}
