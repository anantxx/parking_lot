package service

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"

	. "github.com/parking_lot/constant"
	"github.com/parking_lot/model"
)

var instance *model.Parking
var once sync.Once

type mockParkingRepository struct{}
type mockSlotRepository struct{}

var parSerObj *ParkingService

func init() {
	parkingRepo := mockParkingRepository{}
	slotRepo := mockSlotRepository{}
	parSerObj = NewParkingService(parkingRepo, slotRepo)
}

func GetInstance() *model.Parking {
	once.Do(func() {
		instance = &model.Parking{}
	})
	return instance
}

func (m mockParkingRepository) GetParking() *model.Parking {
	return GetInstance()
}

func (m mockParkingRepository) InitializeLot(totalSlot int) error {
	m.GetParking().TotalSlot = totalSlot
	return nil
}

func (s mockSlotRepository) AddNewSlot(currSlot *model.Slot, newSlot *model.Slot) int {
	//	m.GetParking().AllocateSlot++
	return 0
}

func (s mockSlotRepository) FindSlotBySlotNo(currentSlot *model.Slot, position int) (slot *model.Slot, err error) {
	fmt.Println(currentSlot)
	return &model.Slot{}, nil
}

func (s mockSlotRepository) ReleaseSlot(slot *model.Slot) (err error) {
	//	m.GetParking().AllocateSlot--
	return nil
}

func (s mockSlotRepository) FindAllSlot(slot *model.Slot) []model.Slot {
	return []model.Slot{
		{Car: &model.Car{RegistrationNo: "KA-01-HH-1234",
			Colour: "White",
		},
			Position: 1},
		{Car: &model.Car{RegistrationNo: "KA-01-BB-0001",
			Colour: "Black",
		},
			Position: 2},
		{Car: &model.Car{RegistrationNo: "KA-01-HH-7777",
			Colour: "Red",
		},
			Position: 3},
	}
}

func (s mockSlotRepository) FindSlotsByFeild(slot *model.Slot, value string, feild string) []model.Slot {
	slots := []model.Slot{
		{Car: &model.Car{RegistrationNo: "KA-01-HH-1234",
			Colour: "White",
		},
			Position: 1},
		{Car: &model.Car{RegistrationNo: "KA-01-HH-9999",
			Colour: "White",
		},
			Position: 2},
		{Car: &model.Car{RegistrationNo: "KA-01-BB-0001",
			Colour: "Black",
		},
			Position: 3},
		{Car: &model.Car{RegistrationNo: "KA-01-P-333",
			Colour: "White",
		},
			Position: 4},
		{Car: &model.Car{RegistrationNo: "KA-01-HH-2701",
			Colour: "Blue",
		},
			Position: 5},
		{Car: &model.Car{RegistrationNo: "KA-01-HH-3141",
			Colour: "Red",
		},
			Position: 6},
	}
	var returnSlot []model.Slot
	for _, slot := range slots {
		if "colour" == feild {
			if value == slot.Car.Colour {
				returnSlot = append(returnSlot, slot)
			}
		} else if "registration_number" == feild {
			if value == slot.Car.RegistrationNo {
				returnSlot = append(returnSlot, slot)
			}
		}
	}
	return returnSlot
}

func TestInitializeLot(t *testing.T) {
	testCases := []struct {
		name        string
		input       int
		expected    interface{}
		expectedErr error
	}{
		{"With Negative value", -1, "", ERR_INVALID_ARGUMENT},
		{"Valid Value", 6, "Created a parking lot with 6 slots", nil},
	}
	var (
		output string
		err    error
	)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			output, err = parSerObj.InitializeLot(test.input)
			if test.expected != output {
				t.Errorf("Output %s is not matched with expected %s", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Output  is not matched with expected ")
				}
			}
		})
	}
}

func TestAllocateSlot(t *testing.T) {
	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setTotalSlot           bool
		setLopping             bool
		regNo                  string
		colour                 string
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot", false, true, false, "KA-01-HH-1234", "White", "", ERR_NO_INITIALIZATION},
		{"With Blanck regNo", true, true, false, "", "White", "", ERR_INVALID_ARGUMENT},
		{"With Blanck colour", true, true, false, "KA-01-HH-1234", "", "", ERR_INVALID_ARGUMENT},
		{"With Blanck valid", true, true, false, "KA-01-HH-1234", "White", "Allocated slot number: 1", nil},
		{"With Blanck valid", true, true, true, "KA-01-HH-1234", "White", "", ERR_PARKING_FULL},
	}
	var (
		output string
		err    error
	)

	for _, test := range testCases {
		const totalSlot = 3
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setTotalSlot {
				parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
			}
			if true == test.setLopping {
				for i := 0; i < totalSlot; i++ {
					output, err = parSerObj.AllocateSlot(test.regNo, test.colour)
				}
			}

			output, err = parSerObj.AllocateSlot(test.regNo, test.colour)
			if test.expected != output {
				t.Errorf("Output %s is not matched with expected %s", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Output  is not matched with expected ")
				}
			}
		})
	}
}

func TestReleaseSlot(t *testing.T) {
	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setTotalSlot           bool
		setLopping             bool
		position               int
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot", false, true, false, 2, "", ERR_NO_INITIALIZATION},
		{"With No Car Parked", true, true, false, 2, "", ERR_NO_CAR_PARKED},
		/*{"With Position as 0", true, true, true, 0, "", ERR_CAR_NOT_FOUND},
		{"With Valid Input", true, true, true, 2, "Slot number 2 is free", nil},*/
	}
	var (
		output string
		err    error
	)
	parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = false
	parSerObj.ParkingRepository.GetParking().AllocatedSlot = 0
	for _, test := range testCases {
		const totalSlot = 3
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setTotalSlot {
				parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
			}
			if true == test.setLopping {
				for i := 0; i < totalSlot; i++ {
					output, err = parSerObj.AllocateSlot("KA-01-HH-2701", "Blue")
				}
			}

			output, err = parSerObj.ReleaseSlot(test.position)
			if test.expected != output {
				t.Errorf("Output `%s` is not matched with expected `%s`", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Error: Output `%s` is not matched with expected `%s`", err.Error(), test.expectedErr.Error())
				}
			}
		})
	}
}

func TestShowStatus(t *testing.T) {

	response := fmt.Sprintf("Slot No.\tRegistration No\tColor")
	response += fmt.Sprintf("\n1\t\tKA-01-HH-1234\tWhite")
	response += fmt.Sprintf("\n2\t\tKA-01-BB-0001\tBlack")
	response += fmt.Sprintf("\n3\t\tKA-01-HH-7777\tRed")

	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setCreateSlots         bool
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot", false, false, "", ERR_NO_INITIALIZATION},
		{"With No Car Parked", true, false, "", ERR_NO_CAR_PARKED},
		{"With Valid ", true, true, response, nil},
	}
	var (
		output string
		err    error
	)
	const totalSlot = 3
	parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = false
	parSerObj.ParkingRepository.GetParking().AllocatedSlot = 0
	parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setCreateSlots {
				parSerObj.ParkingRepository.GetParking().AllocatedSlot = 3
			}

			output, err = parSerObj.ShowStatus()
			if test.expected != output {
				t.Errorf("Output `%s` is not matched with expected `%s`", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Error: Output `%s` is not matched with expected `%s`", err.Error(), test.expectedErr.Error())
				}
			}
		})
	}
}

func TestFindRegistationNosByColour(t *testing.T) {

	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setCreateSlots         bool
		colour                 string
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot", false, false, "White", "", ERR_NO_INITIALIZATION},
		{"With No Car Parked", true, false, "White", "", ERR_NO_CAR_PARKED},
		{"With Valid ", true, true, "White", "KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333", nil},
	}
	var (
		output string
		err    error
	)
	const totalSlot = 3
	parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = false
	parSerObj.ParkingRepository.GetParking().AllocatedSlot = 0
	parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setCreateSlots {
				parSerObj.ParkingRepository.GetParking().AllocatedSlot = 3
			}

			output, err = parSerObj.FindRegistationNosByColour(test.colour)
			if test.expected != output {
				t.Errorf("Output `%s` is not matched with expected `%s`", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Error: Output `%s` is not matched with expected `%s`", err.Error(), test.expectedErr.Error())
				}
			}
		})
	}
}

func TestFindAllocatedSlotByColour(t *testing.T) {

	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setCreateSlots         bool
		colour                 string
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot", false, false, "White", "", ERR_NO_INITIALIZATION},
		{"With No Car Parked", true, false, "White", "", ERR_NO_CAR_PARKED},
		{"With Valid ", true, true, "White", "1, 2, 4", nil},
	}
	var (
		output string
		err    error
	)
	const totalSlot = 3
	parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = false
	parSerObj.ParkingRepository.GetParking().AllocatedSlot = 0
	parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setCreateSlots {
				parSerObj.ParkingRepository.GetParking().AllocatedSlot = 3
			}

			output, err = parSerObj.FindAllocatedSlotByColour(test.colour)
			if test.expected != output {
				t.Errorf("Output `%s` is not matched with expected `%s`", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Error: Output `%s` is not matched with expected `%s`", err.Error(), test.expectedErr.Error())
				}
			}
		})
	}
}

func TestFindSlotByRegistationNo(t *testing.T) {

	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setCreateSlots         bool
		regNo                  string
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot", false, false, "KA-01-HH-3141", "", ERR_NO_INITIALIZATION},
		{"With No Car Parked", true, false, "KA-01-HH-3141", "", ERR_NO_CAR_PARKED},
		{"With Valid ", true, true, "KA-01-HH-3141", "6", nil},
		{"With Not Found ", true, true, "MH-04-AY-1111", "", ERR_CAR_NOT_FOUND},
	}
	var (
		output string
		err    error
	)
	const totalSlot = 3
	parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = false
	parSerObj.ParkingRepository.GetParking().AllocatedSlot = 0
	parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setCreateSlots {
				parSerObj.ParkingRepository.GetParking().AllocatedSlot = 3
			}

			output, err = parSerObj.FindSlotByRegistationNo(test.regNo)
			if test.expected != output {
				t.Errorf("Output `%s` is not matched with expected `%s`", output, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Error: Output `%s` is not matched with expected `%s`", err.Error(), test.expectedErr.Error())
				}
			}
		})
	}
}

func TestFindSlotByFeild(t *testing.T) {

	testCases := []struct {
		name                   string
		setIsParkingLotCreated bool
		setCreateSlots         bool
		value                  string
		feild                  string
		expected               string
		expectedErr            error
	}{
		{"Without Creating ParkingLot For Registration Number", false, false, "KA-01-HH-3141", "registration_number", "", ERR_NO_INITIALIZATION},
		{"With No Car Parked For Registration Number", true, false, "KA-01-HH-3141", "registration_number", "", ERR_NO_CAR_PARKED},
		{"With Valid For Registration Number", true, true, "KA-01-HH-3141", "registration_number", "6", nil},
		{"With Not Found For Registration Number", true, true, "MH-04-AY-1111", "registration_number", "", ERR_CAR_NOT_FOUND},
		{"With Not Found For Colour", true, true, "Orange", "colour", "", ERR_CAR_NOT_FOUND},
		{"With Valid For Colour", true, true, "Black", "colour", "3", nil},
		{"With Valid For Colour", true, true, "White", "colour", "1, 2, 4", nil},
	}

	var (
		output []model.Slot
		err    error
	)
	const totalSlot = 3
	parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = false
	parSerObj.ParkingRepository.GetParking().AllocatedSlot = 0
	parSerObj.ParkingRepository.GetParking().TotalSlot = totalSlot
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			t.Helper()
			if true == test.setIsParkingLotCreated {
				parSerObj.ParkingRepository.GetParking().IsParkingLotCreated = true
			}

			if true == test.setCreateSlots {
				parSerObj.ParkingRepository.GetParking().AllocatedSlot = 3
			}

			output, err = parSerObj.FindSlotByFeild(test.value, test.feild)
			position := ""
			for _, s := range output {
				if "" == strings.Trim(position, " ") {
					position = strconv.Itoa(s.Position)
				} else {
					position = position + ", " + strconv.Itoa(s.Position)
				}
			}
			if test.expected != position {
				t.Errorf("Output `%s` is not matched with expected `%s`", position, test.expected)
			}
			if err != nil && test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Errorf("Error: Output `%s` is not matched with expected `%s`", err.Error(), test.expectedErr.Error())
				}
			}
		})
	}
}
