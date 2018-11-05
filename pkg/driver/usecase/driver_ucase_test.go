package usecase

import (
	"github.com/stretchr/testify/assert"
	"gojek-1st/pkg/driver"
	"gojek-1st/pkg/driver/mocks"
	"gojek-1st/pkg/entity"
	"testing"
)

const (
	testLat  = 3.601034
	testLong = 98.679148
)

func Test_DriverService(t *testing.T) {
	t.Run("GetDriverWithinLatLngBounds", GetDriverWithinLatLngBounds_should_get_correct_info)
	t.Run("UpdateDriver_correct_latLng_correct_id", UpdateDriver_correct_latlng_correct_id)
	t.Run("UpdateDriver_incorrect_latlng", UpdateDriver_incorrect_lat_lng)
	t.Run("UpdateDriver_incorrect_id_0", UpdateDriver_incorrect_id_0)
	t.Run("UpdateDriver_incorrect_id_50001", UpdateDriver_incorrect_id_50001)
}

func CreateTestDriver1() entity.Driver {
	test_driver_1 := entity.Driver{
		Accuracy: 0.7,
		Lat:      3.607715,
		Long:     98.672960,
		Id:       2,
	}
	return test_driver_1
}

func CreateTestDriver2() entity.Driver {
	return entity.Driver{
		Accuracy: 0.8,
		Lat:      3.602576,
		Long:     98.681196,
		Id:       3,
	}
}

func CreateTestDriver3() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      3.606002,
		Long:     98.843245,
		Id:       4,
	}
}

func CreateIncorrectLatLngDriver() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      93.606002,
		Long:     98.843245,
		Id:       4,
	}
}

func Create50001IdDriver() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      3.606002,
		Long:     8.843245,
		Id:       50001,
	}
}

func Create0IdDriver() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      3.606002,
		Long:     8.843245,
		Id:       0,
	}
}

func GetDriverWithinLatLngBounds_should_get_correct_info(t *testing.T) {

	driver1 := CreateTestDriver1()
	driver2 := CreateTestDriver2()
	driver3 := CreateTestDriver3()
	mockRepo := new(mocks.Repository)
	allDrivers := []*entity.Driver{&driver1, &driver2, &driver3}
	expectedDrivers := []*entity.Driver{&driver1, &driver2}
	mockRepo.On("GetAll").Return(allDrivers, nil)
	driverUcase := NewDriverUsecase(mockRepo)
	actualDrivers, err := driverUcase.FindDrivers(testLat, testLong, driver.SearchParams{Radius: 1200})
	if err != nil {
		t.Errorf("Unable to find drivers: %s", err.Error())
	}
	assert.ElementsMatch(t, expectedDrivers, actualDrivers, "Expected drivers and the actual drivers are "+
		"not the same!")

}

func UpdateDriver_correct_latlng_correct_id(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectLatLngDriver := CreateIncorrectLatLngDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	err := driverUcase.UpdateLocation(incorrectLatLngDriver.Id, incorrectLatLngDriver.Lat, incorrectLatLngDriver.Long, incorrectLatLngDriver.Accuracy)
	if err != nil {
		t.Error("There is an error!")
	}

}

func UpdateDriver_incorrect_lat_lng(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectLatLngDriver := CreateIncorrectLatLngDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	err := driverUcase.UpdateLocation(incorrectLatLngDriver.Id, incorrectLatLngDriver.Lat, incorrectLatLngDriver.Long, incorrectLatLngDriver.Accuracy)
	if _, ok := err.(*LatLngErr); !ok {
		t.Error("Not LatLngErr")
	}
}

func UpdateDriver_incorrect_id_50001(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectIdDriver := Create50001IdDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	err := driverUcase.UpdateLocation(incorrectIdDriver.Id, incorrectIdDriver.Lat, incorrectIdDriver.Long, incorrectIdDriver.Accuracy)
	if _, ok := err.(*IdErr); !ok {
		t.Error("Not IdErr")
	}

}

func UpdateDriver_incorrect_id_0(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectIdDriver := Create0IdDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	err := driverUcase.UpdateLocation(incorrectIdDriver.Id, incorrectIdDriver.Lat, incorrectIdDriver.Long, incorrectIdDriver.Accuracy)
	if _, ok := err.(*IdErr); !ok {
		t.Error("Not IdErr")
	}

}
