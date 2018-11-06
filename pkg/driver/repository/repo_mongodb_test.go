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

func TestDriverRepo(t *testing.T) {
	t.Run("CreateDriver", createDriver_shouldInsertCorrectly)
	t.Run("UpdateDriver", createDriverAndUpdate_ShouldUpdateCorrectly)
	t.Run("GetDriver", getDriver_shouldGetCorrectDriverInfo)
	t.Run("GetAllDrivers", getAllDrivers_shouldHaveCorrectCount)
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

func createDriver_shouldInsertCorrectly(t *testing.T) {
	readClient, err := CreateMongoClient(mongoUrl)
	writeClient, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(readClient, writeClient, dbName)
	DropDatabase(readClient, dbName)
	driver := CreateTestDriver1()

	_, err = mongoRepository.Store(&driver)

	if err != nil {
		t.Errorf("Unable to create driver: %s", err.Error())
	}
	var results []entity.Driver
	cursor, err := readClient.Database(dbName).Collection(collectionName).Find(context.Background(), nil)

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
		DropDatabase(writeClient, dbName)
	}()
}

func createDriverAndUpdate_ShouldUpdateCorrectly(t *testing.T) {
	readClient, err := CreateMongoClient(mongoUrl)
	writeClient, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(readClient, writeClient, dbName)
	driver := CreateTestDriver1()
	updatedDriver := UpdatedTestDriver1()

	_, err = mongoRepository.Store(&driver)
	_, err = mongoRepository.Store(&updatedDriver)

	if err != nil {
		t.Errorf("Unable to create driver: %s", err.Error())
	}
	var results []entity.Driver
	cursor, err := writeClient.Database(dbName).Collection(collectionName).Find(context.Background(), nil)

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
		DropDatabase(writeClient, dbName)
	}()
}

func getDriver_shouldGetCorrectDriverInfo(t *testing.T) {
	readClient, err := CreateMongoClient(mongoUrl)
	writeClient, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(readClient, writeClient, dbName)
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
		DropDatabase(readClient, dbName)
	}()
}

func getAllDrivers_shouldHaveCorrectCount(t *testing.T) {
	readClient, err := CreateMongoClient(mongoUrl)
	writeClient, err := CreateMongoClient(mongoUrl)
	mongoRepository := CreateMongoRepository(readClient, writeClient, dbName)
	driver1 := CreateTestDriver1()
	driver2 := CreateTestDriver2()
	driver3 := CreateTestDriver3()
	drivers := []*entity.Driver{&driver1, &driver2, &driver3}
	err = mongoRepository.StoreMany(drivers)
	if err != nil {
		t.Errorf("Unable to create drivers: %s", err.Error())
	}
	results, err := mongoRepository.GetAll()
	count := len(results)
	if count != 3 {
		t.Error("Incorrect number of results. Expected `3`, got: `%i`", count)
	}
}
