package repository

import (
	"testing"

	"github.com/parking_lot/constant"
	"github.com/parking_lot/model"
)

var s SlotRepository
var p ParkingRepository

var newSlot *model.Slot
var newSlot1 *model.Slot

func init() {
	newSlot1 = NewSlot(NewCar("KA-01-HH-1234", "White"), 1)
	newSlot2 := NewSlot(NewCar("KA-01-HH-9999", "White"), 2)
	newSlot3 := NewSlot(NewCar("KA-01-BB-0001", "Black"), 3)
	newSlot4 := NewSlot(NewCar("KA-01-HH-7777", "Red"), 4)
	newSlot1.NextSlot = newSlot2
	newSlot2.NextSlot = newSlot3
	newSlot3.NextSlot = newSlot4
	newSlot4.NextSlot = nil

	newSlot = NewSlot(NewCar("KA-01-HH-7777", "Red"), 0)
}

func TestNewSlot(t *testing.T) {
	testCases := []struct {
		name     string
		car      *model.Car
		position int
		expected interface{}
	}{
		{"With Postion 2", NewCar("KA-01-HH-1234", "White"), 2, "KA-01-HH-1234"},
		{"With Postion 10", NewCar("KA-01-HH-999", "Blue"), 10, "KA-01-HH-999"},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			slot := NewSlot(test.car, test.position)
			if test.expected != slot.Car.RegistrationNo {
				t.Errorf("Output %s is not matched with expected %s", slot.Car.RegistrationNo, test.expected)
			}
		})
	}
}
func TestAddNewSlot(t *testing.T) {
	testCases := []struct {
		name        string
		currentSlot *model.Slot
		newSlot     *model.Slot
		expected    interface{}
	}{
		{"With Empty Slot", newSlot1, newSlot, 5},
		{"With Next Slot", newSlot1, newSlot, 6},
	}
	var (
		pos int
	)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			pos = s.AddNewSlot(test.currentSlot, test.newSlot)
			if test.expected != pos {
				t.Errorf("Output %d is not matched with expected %d", pos, test.expected)
			}
		})
	}
}

func TestFindSlotBySlotNo(t *testing.T) {
	testCases := []struct {
		name        string
		currentSlot *model.Slot
		position    int
		expected    interface{}
		expectedErr error
	}{
		{"With Valid Postion", newSlot1, 3, "KA-01-BB-0001", nil},
		{"With Invalid Position", newSlot1, 7, nil, constant.ERR_CAR_NOT_FOUND},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			slot, err := s.FindSlotBySlotNo(test.currentSlot, test.position)
			if nil != slot {
				if test.expected != slot.Car.RegistrationNo {
					t.Errorf("Output %s is not matched with expected %s", slot.Car.RegistrationNo, test.expected)
				}
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Output is not matched with expected ")
				}
			}
		})
	}
}

func TestFindAllSlot(t *testing.T) {
	testCases := []struct {
		name        string
		currentSlot *model.Slot
		expected    interface{}
	}{
		{"With Valid Postion", newSlot1, 5},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			slot := s.FindAllSlot(test.currentSlot)
			if test.expected != len(slot) {
				t.Errorf("Output %d is not matched with expected %d", len(slot), test.expected)
			}
		})
	}
}

func TestFindSlotsByFeild(t *testing.T) {
	testCases := []struct {
		name        string
		currentSlot *model.Slot
		value       string
		feild       string
		expected    interface{}
	}{
		{"With Valid Colour", newSlot1, "White", "colour", 2},
		{"With Valid Registration Number", newSlot1, "KA-01-HH-1234", "registration_number", 1},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			slot := s.FindSlotsByFeild(test.currentSlot, test.value, test.feild)
			if test.expected != len(slot) {
				t.Errorf("Output %d is not matched with expected %d", len(slot), test.expected)
			}
		})
	}
}

func TestReleaseSlot(t *testing.T) {
	testCases := []struct {
		name        string
		currentSlot *model.Slot
		position    int
		expected    interface{}
		expectedErr error
	}{
		{"With Valid Postion", newSlot1, 3, "KA-01-BB-0001", nil},
		{"With Invalid Position", newSlot1, 7, nil, constant.ERR_CAR_NOT_FOUND},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			slot, err := s.FindSlotBySlotNo(test.currentSlot, test.position)
			if err == nil {
				err = s.ReleaseSlot(slot)
			}
			if nil != slot {
				if test.expected != slot.Car.RegistrationNo {
					t.Errorf("Output %s is not matched with expected %s", slot.Car.RegistrationNo, test.expected)
				}
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Output is not matched with expected ")
				}
			}
		})
	}
}
