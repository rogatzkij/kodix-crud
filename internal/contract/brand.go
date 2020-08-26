package contract

import "github.com/rogatzkij/kodix-crud/model"

type Brand interface {
	CreateBrand(brand model.Brand) error
	DeleteBrand(brandname string) error
	CheckBrand(brandname string) (bool, error)
	CreateModel(brandname, automodel string) error
	CheckModel(brandname, automodel string) (bool, error)
	DeleteModel(brandname, automodel string) error
}
