package driver

import "gojek-1st/pkg/entity"

type Repository interface {
	Store(d *entity.Driver) (status string, err error)
	Get(id string) (d *entity.Driver, err error)
}
