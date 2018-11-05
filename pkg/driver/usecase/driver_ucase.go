package usecase

import (
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

func (du *DriverUsecase) FindDrivers(latitude float64, longitude float64, radius float64, limit int8) (drivers []*entity.Driver, err error) {
	return nil, nil
}

func (du *DriverUsecase) UpdateLocation(id string) (err error) {
	return nil
}
