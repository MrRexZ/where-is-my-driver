package usecase

import (
	"fmt"
	"github.com/umahmood/haversine"
	"gojek-1st/pkg/driver"
	"gojek-1st/pkg/entity"
	"math"
)

const (
	lowestId  = 1
	highestId = 50000
)

type LatLngErr struct {
	Msg string
}

type IdErr struct {
	Msg string
}

func (e *LatLngErr) Error() string {
	return fmt.Sprintf("LatLng error : %s", e.Msg)
}

func (e *IdErr) Error() string {
	return fmt.Sprintf("Id error : %s", e.Msg)
}

type DriverUsecase struct {
	repo driver.Repository
}

func NewDriverUsecase(repo driver.Repository) *DriverUsecase {
	return &DriverUsecase{
		repo: repo,
	}
}

func (du *DriverUsecase) FindDrivers(latitude float64, longitude float64, radius float64, limit int8) (drivers []*entity.Driver, err error) {

	var count int8 = 0
	var driversFiltered []*entity.Driver
	drivers, err = du.repo.GetAll()
	for _, driver := range drivers {
		if getMetersDistanceLatsLngs(latitude, longitude, driver.Lat, driver.Long) <= radius {
			driversFiltered = append(driversFiltered, driver)
			count += 1
		}
		if count == limit {
			break
		}
	}
	return driversFiltered, nil
}

func getMetersDistanceLatsLngs(la1 float64, lo1 float64, la2 float64, lo2 float64) float64 {
	coord1 := haversine.Coord{la1, lo1}
	coord2 := haversine.Coord{la2, lo2}
	_, metersDistance := haversine.Distance(coord1, coord2)
	return metersDistance * 1000

}

func (du *DriverUsecase) UpdateLocation(id int32, lat float64, long float64, accuracy float64) (int32, error) {
	if !du.IsValidId(id) {
		return 0, &IdErr{"ID out of bound"}
	}
	if !du.IsValidLatLng(lat, long) {
		return 0, &LatLngErr{"LatLng should be between +/-90"}
	}
	driver := entity.Driver{Id: id, Lat: lat, Long: long, Accuracy: accuracy}
	id, err := du.repo.Store(&driver)
	return id, err
}

func (du *DriverUsecase) IsValidLatLng(lat float64, long float64) (valid bool) {
	if math.Abs(lat) > 90 || math.Abs(long) > 90 {
		return false
	}
	return true
}

func (du *DriverUsecase) IsValidId(id int32) (valid bool) {
	if id >= lowestId && id <= highestId {
		return true
	}
	return false
}
