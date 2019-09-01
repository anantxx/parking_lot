package repository

import "github.com/parking_lot/model"

// NewCar returns a new car object
func NewSlot(car *model.Car, pos int) *model.Slot {
	return &model.Slot{
		Car:      car,
		Position: pos,
	}
}
