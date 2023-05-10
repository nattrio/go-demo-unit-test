package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amount must be greater than zero")
	ErrRepository = errors.New("unable to get promotion from repository")
)
