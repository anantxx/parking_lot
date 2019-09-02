package repository

import (
	"github.com/parking_lot/constant"
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
		pos := currSlot.Position + 1
		newSlot.Position = pos
		currSlot.NextSlot = newSlot
		newSlot.PrevSlot = currSlot
		return pos
	}

	if currSlot.NextSlot.Position > (currSlot.Position + 1) {
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

func (s SlotRepository) FindSlotBySlotNo(currentSlot *model.Slot, position int) (slot *model.Slot, err error) {
	if position == currentSlot.Position {
		return currentSlot, nil
	}

	if nil == currentSlot.NextSlot {
		return nil, constant.ERR_CAR_NOT_FOUND
	}

	return s.FindSlotBySlotNo(currentSlot.NextSlot, position)
}

func (s SlotRepository) ReleaseSlot(slot *model.Slot) (err error) {
	if slot.PrevSlot != nil {
		slot.PrevSlot.NextSlot = slot.NextSlot
	}
	if slot.NextSlot != nil {
		slot.NextSlot.PrevSlot = slot.PrevSlot
	}
	slot = nil
	return err
}

func (s SlotRepository) FindAllSlot(slot *model.Slot) []model.Slot {
	if slot == nil {
		return []model.Slot{}
	}
	return append([]model.Slot{*slot}, s.FindAllSlot(slot.NextSlot)...)
}

func (s SlotRepository) FindSlotsByFeild(slot *model.Slot, value string, feild string) []model.Slot {
	if slot == nil {
		return []model.Slot{}
	}
	if "colour" == feild {
		if value == slot.Car.Colour {
			return append([]model.Slot{*slot}, s.FindSlotsByFeild(slot.NextSlot, value, feild)...)
		}
	} else if "registration_number" == feild {
		if value == slot.Car.RegistrationNo {
			return append([]model.Slot{*slot}, s.FindSlotsByFeild(slot.NextSlot, value, feild)...)
		}
	}

	return s.FindSlotsByFeild(slot.NextSlot, value, feild)
}
