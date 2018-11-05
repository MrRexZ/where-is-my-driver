package driver

import "gojek-1st/pkg/entity"

type Repository interface {
	StoreMany(ds []*entity.Driver) error
	Store(d *entity.Driver) (id string, err error)
	Get(id string) (d *entity.Driver, err error)
	GetAll() (ds []*entity.Driver, err error)
}
