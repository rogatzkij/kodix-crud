package contract

import "github.com/rogatzkij/kodix-crud/model"

type Auto interface {
	Create(auto model.Auto) (uint, error)
	GetByID(id uint) (model.Auto, error)
	UpdateByID(id uint, auto model.Auto) error
	DeleteByID(id uint) error
}
