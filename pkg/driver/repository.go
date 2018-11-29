package driver

import "where-is-my-driver/pkg/entity"

type Repository interface {
	StoreMany(ds []*entity.Driver) error
	Store(d *entity.Driver) (id int32, err error)
	Get(id int32) (d *entity.Driver, err error)
	GetAll() (ds []*entity.Driver, err error)
}
