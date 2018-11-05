package usecase

import (
	"github.com/umahmood/haversine"
	"gojek-1st/pkg/driver"
	"gojek-1st/pkg/entity"
)

type DriverUsecase struct {
	repo driver.Repository
}

func NewDriverUsecase(repo driver.Repository) *DriverUsecase {
	return &DriverUsecase{
		repo: repo,
	}
}

func (du *DriverUsecase) FindDrivers(latitude float64, longitude float64, params driver.SearchParams) (drivers []*entity.Driver, err error) {
	var radius float64 = 500
	var limit int8 = 10
	var count int8 = 0
	var driversFiltered []*entity.Driver
	if params.Radius != 0 {
		radius = params.Radius
	}
	if params.Limit != 0 {
		limit = params.Limit
	}
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

func (du *DriverUsecase) UpdateLocation(id string) (err error) {
	return nil
}
