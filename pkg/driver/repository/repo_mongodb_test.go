package repository

import (
	"context"
	"gojek-1st/pkg/entity"
	"log"
	"testing"
)

const (
	mongoUrl = "mongodb://localhost:27017"
	dbName   = "test_db"
)

func Test_DriverService(t *testing.T) {
	t.Run("CreateDriver", createDriver_should_insert_correctly)
}

func createDriver_should_insert_correctly(t *testing.T) {
	client, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(client, dbName)
	driver := entity.Driver{
		Accuracy: 0.7,
		Lat:      12.3,
		Long:     23.1,
		Id:       2,
	}

	_, err = mongoRepository.Store(&driver)

	if err != nil {
		t.Errorf("Unable to create driver: %s", err.Error())
	}
	var results []entity.Driver
	cursor, err := client.Database(dbName).Collection(collectionName).Find(context.Background(), nil)

	for cursor.Next(context.Background()) {
		driver := entity.Driver{}
		err := cursor.Decode(&driver)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, driver)

	}
	count := len(results)
	if count != 1 {
		t.Error("Incorrect number of results. Expected `1`, got: `%i`", count)
	}
	if results[0].Id != driver.Id {
		t.Error("Incorrect Id. Expected `%i`, Got: `%i`", driver.Id, results[0].Id)
	}

	defer func() {
		DropDatabase(client, dbName)
	}()
}