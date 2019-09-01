package repository

import "github.com/parking_lot/model"

// NewCar returns a new car object
func NewCar(regNo, colour string) *model.Car {
	return &model.Car{
		RegistrationNo: regNo,
		Colour:         colour,
	}
}
