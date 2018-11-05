package driver

import "gojek-1st/pkg/entity"

type Usecase interface {
	UpdateLocation(id string) (err error)
	FindDrivers(latitude float64, longitude float64, radius float64, limit int8) (drivers []*entity.Driver, err error)
}
