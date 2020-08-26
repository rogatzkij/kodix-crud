package contract

import "github.com/rogatzkij/kodix-crud/model"

type Brand interface {
	CreateBrand(brand model.Brand) error
	CheckBrand(brand string) (bool, error)
	CreateModel(brand, model string) error
	CheckModel(brand, model string) (bool, error)
}
