package repository

import (
	"fmt"

	"github.com/parking_lot/model"
)

type SlotRepository struct {
	Slot *model.Slot
}

func NewSlotRepository() SlotRepository {
	return SlotRepository{}
}

// NewCar returns a new car object
func NewSlot(car *model.Car, pos int) *model.Slot {
	return &model.Slot{
		Car:      car,
		Position: pos,
	}
}

func (s SlotRepository) AddNewSlot(currSlot *model.Slot, newSlot *model.Slot) int {
	if nil == currSlot.NextSlot {
		fmt.Println("INISDE FIRST 1")
		pos := currSlot.Position + 1
		newSlot.Position = pos
		currSlot.NextSlot = newSlot
		newSlot.PrevSlot = currSlot
		return pos
	}

	if currSlot.NextSlot.Position > (currSlot.Position + 1) {
		fmt.Println("INISDE FIRST 2")
		pos := currSlot.Position + 1
		newSlot.Position = pos
		currentNext := currSlot.NextSlot
		currSlot.NextSlot = newSlot
		newSlot.PrevSlot = currSlot
		newSlot.NextSlot = currentNext
		currentNext.PrevSlot = newSlot
		return pos
	}
	return s.AddNewSlot(currSlot.NextSlot, newSlot)
}
