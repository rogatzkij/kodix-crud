package contract

import "github.com/rogatzkij/kodix-crud/model"

type Auto interface {
	CreateAuto(auto model.Auto) (uint, error)
	GetAutos(limit, offset uint) ([]model.Auto, error)
	GetAutoByID(id uint) (*model.Auto, error)
	UpdateAutoByID(id uint, auto model.Auto) error
	DeleteAutoByID(id uint) error
}
