package usecase

import (
	"github.com/stretchr/testify/assert"
	"gojek-1st/pkg/driver/mocks"
	"gojek-1st/pkg/entity"
	"testing"
)

const (
	testLat  = 51.507351
	testLong = -0.127758
)

func Test_DriverService(t *testing.T) {
	t.Run("GetDriverWithinLatLngBounds", GetDriverWithinLatLngBounds_should_get_correct_info)
	t.Run("IsValidLatLng_valid", IsValidLatLng_valid)
	t.Run("IsValidLatLng_invalid", IsValidLatLng_invalid)
	t.Run("IsValidId_valid", IsValidId_valid)
	t.Run("IsValidId_invalid", IsValidId_invalid)
	t.Run("UpdateDriver_correct_latLng_correct_id", UpdateDriver_correct_latlng_correct_id)
	t.Run("UpdateDriver_incorrect_latlng", UpdateDriver_incorrect_lat_lng)
	t.Run("UpdateDriver_incorrect_id_0", UpdateDriver_incorrect_id_0)
	t.Run("UpdateDriver_incorrect_id_50001", UpdateDriver_incorrect_id_50001)
}

func CreateTestDriver1() entity.Driver {
	test_driver_1 := entity.Driver{
		Accuracy: 0.7,
		Lat:      51.506752,
		Long:     -0.132912,
		Id:       2,
	}
	return test_driver_1
}

func CreateTestDriver2() entity.Driver {
	return entity.Driver{
		Accuracy: 0.8,
		Lat:      51.508888,
		Long:     -0.125706,
		Id:       3,
	}
}

func CreateTestDriver3() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      51.508034,
		Long:     0.084315,
		Id:       4,
	}
}

func CreateValidDriver() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      3.606002,
		Long:     8.843245,
		Id:       4,
	}
}

func CreateIncorrectLatLngDriver() entity.Driver {
	lat, lng := InvalidLatLng()
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      lat,
		Long:     lng,
		Id:       4,
	}
}

func CreateUpperboundIdDriver() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      3.606002,
		Long:     8.843245,
		Id:       highestId + 1,
	}
}

func CreateLowerboundIdDriver() entity.Driver {
	return entity.Driver{
		Accuracy: 0.7,
		Lat:      3.606002,
		Long:     8.843245,
		Id:       lowestId - 1,
	}
}

func ValidLatLng() (lat float64, long float64) {
	return 3, 2
}

func ValidLatInvalidLng() (lat float64, long float64) {
	return 89, 91
}

func InvalidLatValidLng() (lat float64, long float64) {
	return 91, 89
}

func InvalidLatLng() (lat float64, long float64) {
	return 91, -91
}

func ValidId() (id int32) {
	return 1
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
	actualDrivers, err := driverUcase.FindDrivers(testLat, testLong, 1200, 10)
	if err != nil {
		t.Errorf("Unable to find drivers: %s", err.Error())
	}
	assert.ElementsMatch(t, expectedDrivers, actualDrivers, "Expected drivers and the actual drivers are "+
		"not the same!")

}

func UpdateDriver_correct_latlng_correct_id(t *testing.T) {
	mockRepo := new(mocks.Repository)
	correctLatLngDriver := CreateValidDriver()
	mockRepo.On("Store", &correctLatLngDriver).Return(correctLatLngDriver.Id, nil)
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(correctLatLngDriver.Id, correctLatLngDriver.Lat, correctLatLngDriver.Long, correctLatLngDriver.Accuracy)
	if err != nil {
		t.Error("There is an error!")
	}

}

func UpdateDriver_incorrect_lat_lng(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectLatLngDriver := CreateIncorrectLatLngDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(incorrectLatLngDriver.Id, incorrectLatLngDriver.Lat, incorrectLatLngDriver.Long, incorrectLatLngDriver.Accuracy)
	if _, ok := err.(*LatLngErr); !ok {
		t.Error("Not LatLngErr")
	}
}

func UpdateDriver_incorrect_id_50001(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectIdDriver := CreateUpperboundIdDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(incorrectIdDriver.Id, incorrectIdDriver.Lat, incorrectIdDriver.Long, incorrectIdDriver.Accuracy)
	if _, ok := err.(*IdErr); !ok {
		t.Error("Not IdErr")
	}

}

func UpdateDriver_incorrect_id_0(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectIdDriver := CreateLowerboundIdDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(incorrectIdDriver.Id, incorrectIdDriver.Lat, incorrectIdDriver.Long, incorrectIdDriver.Accuracy)
	if _, ok := err.(*IdErr); !ok {
		t.Error("Not IdErr")
	}

}

func IsValidLatLng_valid(t *testing.T) {

	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.True(t, driverUcase.IsValidLatLng(ValidLatLng()))
}

func IsValidLatLng_invalid(t *testing.T) {

	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.False(t, driverUcase.IsValidLatLng(InvalidLatLng()))
	assert.False(t, driverUcase.IsValidLatLng(InvalidLatValidLng()))
	assert.False(t, driverUcase.IsValidLatLng(ValidLatInvalidLng()))
}

func IsValidId_valid(t *testing.T) {
	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.True(t, driverUcase.IsValidId(ValidId()))
}

func IsValidId_invalid(t *testing.T) {
	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.False(t, driverUcase.IsValidId(highestId+1))
	assert.False(t, driverUcase.IsValidId(lowestId-1))
}
