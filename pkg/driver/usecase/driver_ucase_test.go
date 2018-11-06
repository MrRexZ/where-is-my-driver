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

func TestDriverUcase(t *testing.T) {
	t.Run("GetDriverWithinLatLngBounds", getDriverWithinLatLngBounds_shouldGetCorrectInfo)
	t.Run("getDriver_invalidLatLng", getDriver_invalidLatLng)
	t.Run("isValidLatLng_valid", isValidLatLng_valid)
	t.Run("isValidLatLng_invalid", isValidLatLng_invalid)
	t.Run("isValidId_valid", isValidId_valid)
	t.Run("isValidId_invalid", isValidId_invalid)
	t.Run("UpdateDriver_correct_latLng_correct_id", updateDriver_correctLatlngCorrectId)
	t.Run("UpdateDriver_incorrect_latlng", updateDriverIncorrectLatLng)
	t.Run("updateDriver_incorrectIdLowerbound", updateDriver_incorrectIdLowerbound)
	t.Run("updateDriver_incorrectIdUpperbound", updateDriver_incorrectIdUpperbound)
}

func createTestDriver1() entity.Driver {
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

func getDriverWithinLatLngBounds_shouldGetCorrectInfo(t *testing.T) {

	driver1 := createTestDriver1()
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

func getDriver_invalidLatLng(t *testing.T) {
	invLat, invLong := InvalidLatLng()
	mockRepo := new(mocks.Repository)
	correctIdDriver := CreateValidDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(correctIdDriver.Id, invLat, invLong, correctIdDriver.Accuracy)
	_, ok := err.(*LatLngErr)
	assert.True(t, ok)
}

func updateDriver_correctLatlngCorrectId(t *testing.T) {
	mockRepo := new(mocks.Repository)
	correctLatLngDriver := CreateValidDriver()
	mockRepo.On("Store", &correctLatLngDriver).Return(correctLatLngDriver.Id, nil)
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(correctLatLngDriver.Id, correctLatLngDriver.Lat, correctLatLngDriver.Long, correctLatLngDriver.Accuracy)
	if err != nil {
		t.Error("There is an error!")
	}

}

func updateDriverIncorrectLatLng(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectLatLngDriver := CreateIncorrectLatLngDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(incorrectLatLngDriver.Id, incorrectLatLngDriver.Lat, incorrectLatLngDriver.Long, incorrectLatLngDriver.Accuracy)
	if _, ok := err.(*LatLngErr); !ok {
		t.Error("Not LatLngErr")
	}
}

func updateDriver_incorrectIdUpperbound(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectIdDriver := CreateUpperboundIdDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(incorrectIdDriver.Id, incorrectIdDriver.Lat, incorrectIdDriver.Long, incorrectIdDriver.Accuracy)
	_, ok := err.(*IdErr)
	assert.True(t, ok)

}

func updateDriver_incorrectIdLowerbound(t *testing.T) {
	mockRepo := new(mocks.Repository)
	incorrectIdDriver := CreateLowerboundIdDriver()
	driverUcase := NewDriverUsecase(mockRepo)
	_, err := driverUcase.UpdateLocation(incorrectIdDriver.Id, incorrectIdDriver.Lat, incorrectIdDriver.Long, incorrectIdDriver.Accuracy)
	_, ok := err.(*IdErr)
	assert.True(t, ok)

}

func isValidLatLng_valid(t *testing.T) {

	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.True(t, driverUcase.IsValidLatLng(ValidLatLng()))
}

func isValidLatLng_invalid(t *testing.T) {

	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.False(t, driverUcase.IsValidLatLng(InvalidLatLng()))
	assert.False(t, driverUcase.IsValidLatLng(InvalidLatValidLng()))
	assert.False(t, driverUcase.IsValidLatLng(ValidLatInvalidLng()))
}

func isValidId_valid(t *testing.T) {
	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.True(t, driverUcase.IsValidId(ValidId()))
}

func isValidId_invalid(t *testing.T) {
	mockRepo := new(mocks.Repository)
	driverUcase := NewDriverUsecase(mockRepo)
	assert.False(t, driverUcase.IsValidId(highestId+1))
	assert.False(t, driverUcase.IsValidId(lowestId-1))
}
