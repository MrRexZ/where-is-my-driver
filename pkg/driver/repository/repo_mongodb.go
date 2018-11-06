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
	readClient  *mongo.Client
	writeClient *mongo.Client
	db          string
}

func NewMongoRepository(readClient *mongo.Client, writeClient *mongo.Client, db string) *MongoRepository {
	return &MongoRepository{
		readClient:  readClient,
		writeClient: writeClient,
		db:          db,
	}
}

func (mr *MongoRepository) Store(d *entity.Driver) (int32, error) {
	_, err := mr.writeClient.Database(mr.db).Collection(collectionName).ReplaceOne(nil, bson.NewDocument(
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
	doc_res := mr.readClient.Database(mr.db).Collection(collectionName).FindOne(nil,
		bson.NewDocument(bson.EC.Int32("id", id)))
	doc_res.Decode(&driver)
	return &driver, err
}

func (mr *MongoRepository) GetAll() (ds []*entity.Driver, err error) {
	cursor, err := mr.readClient.Database(mr.db).Collection(collectionName).Find(nil, nil)
	var drivers []*entity.Driver
	for cursor.Next(context.Background()) {
		driver := entity.Driver{}
		err = cursor.Decode(&driver)
		drivers = append(drivers, &driver)
	}
	return drivers, err
}
