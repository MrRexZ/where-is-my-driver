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
