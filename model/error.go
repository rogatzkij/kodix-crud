package model

import "errors"

var (
	ErrBrandAlreadyExist       = errors.New("the brand already exist")
	ErrBrandDoesntAlreadyExist = errors.New("the brand doesn't exist")
)
