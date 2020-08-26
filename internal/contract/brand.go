package contract

import "github.com/rogatzkij/kodix-crud/model"

type Brand interface {
	CreateBrand(brand model.Brand) error
	DeleteBrand(brandname string) error
	CheckBrand(brandname string) (bool, error)
	CreateModel(brandname, model string) error
	CheckModel(brandname, model string) (bool, error)
	DeleteModel(brandname, model string) error
}
