package model

import "errors"

var (
	ErrBrandAlreadyExist       = errors.New("the brand already exist")
	ErrBrandDoesntAlreadyExist = errors.New("the brand doesn't exist")

	ErrAutomodelAlreadyExist       = errors.New("the model already exist")
	ErrAutomodelDoesntAlreadyExist = errors.New("the model doesn't exist")
)
