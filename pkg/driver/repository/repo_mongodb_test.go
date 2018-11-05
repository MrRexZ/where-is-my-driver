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
	t.Run("UpdateDriver", createDriverAndUpdate_should_update_correctly)
	t.Run("GetDriver", GetDriver_should_get_correct_driver_info)
}

func CreateTestDriver1() entity.Driver {
	test_driver_1 := entity.Driver{
		Accuracy: 0.7,
		Lat:      12.3,
		Long:     23.1,
		Id:       2,
	}
	return test_driver_1
}

func UpdatedTestDriver1() entity.Driver {
	test_driver_1 := entity.Driver{
		Accuracy: 0.6,
		Lat:      2.3,
		Long:     1.1,
		Id:       2,
	}
	return test_driver_1
}

func CreateTestDriver2() entity.Driver {
	return entity.Driver{
		Accuracy: 0.8,
		Lat:      1.3,
		Long:     1.1,
		Id:       3,
	}
}

func CreateTestDriver3() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      19.3,
		Long:     10.1,
		Id:       4,
	}
}

func createDriver_should_insert_correctly(t *testing.T) {
	client, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(client, dbName)
	driver := CreateTestDriver1()

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

func createDriverAndUpdate_should_update_correctly(t *testing.T) {
	client, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(client, dbName)
	driver := CreateTestDriver1()
	updatedDriver := UpdatedTestDriver1()

	_, err = mongoRepository.Store(&driver)
	_, err = mongoRepository.Store(&updatedDriver)

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
	if results[0].Id != updatedDriver.Id {
		t.Error("Incorrect Id. Expected `%i`, Got: `%i`", driver.Id, results[0].Id)
	}
	if results[0].Lat != updatedDriver.Lat || results[0].Long != updatedDriver.Long {
		t.Error("Incorrect LatLong. Expected `%i`, `%i`, Got: `%i`, `%i`", updatedDriver.Lat, updatedDriver.Long, results[0].Lat, results[0].Long)
	}

	defer func() {
		DropDatabase(client, dbName)
	}()
}

func GetDriver_should_get_correct_driver_info(t *testing.T) {
	client, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(client, dbName)
	driver := CreateTestDriver1()

	_, err = mongoRepository.Store(&driver)

	if err != nil {
		t.Errorf("Unable to create driver: %s", err.Error())
	}
	actual_driver, err := mongoRepository.Get(driver.Id)

	if actual_driver.Id != driver.Id {
		t.Error("Incorrect Id. Expected `%i`, Got: `%i`", driver.Id, actual_driver.Id)
	}
	if actual_driver.Lat != driver.Lat || actual_driver.Long != driver.Long {
		t.Error("Incorrect LatLong. Expected `%i`, `%i`, Got: `%i`, `%i`", driver.Lat, driver.Long, actual_driver.Lat, actual_driver.Long)
	}

	defer func() {
		DropDatabase(client, dbName)
	}()
}
