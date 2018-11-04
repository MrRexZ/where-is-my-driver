package driver

import "gojek-1st/pkg/entity"

type Usecase interface {
	UpdateLocation(id string) (status string, err error)
	FindDrivers(latitude float64, longitude float64, radius float64, limit int8) (status string, err error, drivers []*entity.Driver)
}
