package repository

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"gojek-1st/pkg/entity"
)

type MongoRepository struct {
	client *mongo.Client
	db     string
}

func NewMongoRepository(client *mongo.Client, db string) *MongoRepository {
	return &MongoRepository{
		client: client,
		db:     db,
	}
}

func (mr *MongoRepository) Store(d *entity.Driver) (string, error) {

}

func (mr *MongoRepository) Get(id string) (d *entity.Driver, err error) {

}
