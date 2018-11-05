package driver

import "gojek-1st/pkg/entity"

type SearchParams struct {
	Radius float64
	Limit  int8
}
type Usecase interface {
	UpdateLocation(id string) (err error)
	FindDrivers(latitude float64, longitude float64, params SearchParams) (drivers []*entity.Driver, err error)
}
