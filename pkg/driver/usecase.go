package driver

import "gojek-1st/pkg/entity"

type SearchParams struct {
	Radius float64
	Limit  int8
}
type Usecase interface {
	UpdateLocation(id int32, lat float64, long float64, accuracy float64) (err error)
	FindDrivers(latitude float64, longitude float64, params SearchParams) (drivers []*entity.Driver, err error)
	IsValidLatLng(latitude float64, longitude float64) error
	IsValidId(id int32)
}
