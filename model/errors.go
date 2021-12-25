package model

import "errors"

var (
	ErrConflict         = errors.New("Object Conflict")
	ErrAlreadyExists    = errors.New("Dialplan for number already exists")
	ErrSteps            = errors.New("Steps are not defined properly")
	ErrDatabase         = errors.New("Database error")
	ErrIncorrectRequest = errors.New("Request format is incorrect")
	ErrTryAgain         = errors.New("Please try again")
	ErrNotExists        = errors.New("Object does not exist")
)

const (
	CodeAlreadyExists    = 1
	CodeSteps            = 2
	CodeDatabase         = 3
	CodeIncorrectRequest = 4
	CodeTryAgain         = 5
	CodeNotExists        = 6
	CodeConflict         = 7
)
