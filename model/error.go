package model

import "errors"

var (
	ErrBrandAlreadyExist = errors.New("the brand already exist")
	ErrBrandDoesntExist  = errors.New("the brand doesn't exist")

	ErrAutomodelAlreadyExist = errors.New("the model already exist")
	ErrAutomodelDoesntExist  = errors.New("the model doesn't exist")

	ErrBrandOrModelDoesntExist = errors.New("the brand or the model doesn't exist")
)
