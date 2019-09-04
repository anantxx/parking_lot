package service

import (
	"fmt"
	"strings"

	. "github.com/parking_lot/constant"
	"github.com/parking_lot/model"
	"github.com/parking_lot/repository"
)

type ParkingService struct {
	ParkingRepository repository.RepositoryInstance
	SlotRepository    repository.SlotRepositoryInstance
}

func NewParkingService(parkingRepo repository.RepositoryInstance, slotRepo repository.SlotRepositoryInstance) *ParkingService {
	return &ParkingService{
		ParkingRepository: parkingRepo,
		SlotRepository:    slotRepo,
	}
}

func (p *ParkingService) InitializeLot(totalSlot int) (string, error) {
	if totalSlot < 0 {
		return "", ERR_INVALID_ARGUMENT
	}

	err := p.ParkingRepository.InitializeLot(totalSlot)
	if nil != err {
		return "", err
	}
	response := fmt.Sprintf("Created a parking lot with %d slots", p.getPaking().TotalSlot)
	return response, nil
}

func (p *ParkingService) getPaking() *model.Parking {
	return p.ParkingRepository.GetParking()
}

func (p *ParkingService) AllocateSlot(regNo string, colour string) (string, error) {

	if false == p.getPaking().IsParkingLotCreated {
		return "", ERR_NO_INITIALIZATION
	}

	if "" == strings.Trim(regNo, " ") || "" == strings.Trim(colour, " ") {
		return "", ERR_INVALID_ARGUMENT
	}

	if p.getPaking().TotalSlot <= p.getPaking().AllocatedSlot {
		return "", ERR_PARKING_FULL
	}

	car := repository.NewCar(regNo, colour)
	var pos int
	if 0 == p.getPaking().AllocatedSlot {
		slot := repository.NewSlot(car, 1)
		p.getPaking().Slots = slot
		pos = slot.Position
	} else {
		newSlot := repository.NewSlot(car, 0)
		pos = p.SlotRepository.AddNewSlot(p.getPaking().Slots, newSlot)
	}
	p.getPaking().AllocatedSlot++
	return fmt.Sprintf("Allocated slot number: %d", pos), nil
}

func (p *ParkingService) ReleaseSlot(slotNo int) (string, error) {
	if false == p.getPaking().IsParkingLotCreated {
		return "", ERR_NO_INITIALIZATION
	}

	if p.getPaking().AllocatedSlot < 1 {
		return "", ERR_NO_CAR_PARKED
	}

	realeaseSlot, err := p.SlotRepository.FindSlotBySlotNo(p.getPaking().Slots, slotNo)
	if nil != err {
		return "", err
	}

	err = p.SlotRepository.ReleaseSlot(realeaseSlot)
	if nil != err {
		return "", err
	}
	p.getPaking().AllocatedSlot--
	return fmt.Sprintf("Slot number %d is free", realeaseSlot.Position), nil

}

func (p *ParkingService) ShowStatus() (string, error) {
	if false == p.getPaking().IsParkingLotCreated {
		return "", ERR_NO_INITIALIZATION
	}
	if p.getPaking().AllocatedSlot < 1 {
		return "", ERR_NO_CAR_PARKED
	}

	slots := p.SlotRepository.FindAllSlot(p.getPaking().Slots)
	response := fmt.Sprintf("Slot No.\tRegistration No\tColor")
	for _, slot := range slots {
		response += fmt.Sprintf("\n%d\t\t%s\t%s", slot.Position, slot.Car.RegistrationNo, slot.Car.Colour)
	}
	return response, nil
}

func (p *ParkingService) FindSlotByFeild(value string, feild string) ([]model.Slot, error) {
	if false == p.getPaking().IsParkingLotCreated {
		return []model.Slot{}, ERR_NO_INITIALIZATION
	}

	if p.getPaking().AllocatedSlot < 1 {
		return []model.Slot{}, ERR_NO_CAR_PARKED
	}

	slots := p.SlotRepository.FindSlotsByFeild(p.getPaking().Slots, value, feild)
	if len(slots) < 1 {
		return []model.Slot{}, ERR_CAR_NOT_FOUND
	}
	return slots, nil
}

func (p *ParkingService) FindRegistationNosByColour(colour string) (string, error) {
	slots, err := p.FindSlotByFeild(colour, "colour")
	if nil != err {
		return "", err
	}
	var response string
	for _, slot := range slots {
		if "" == response {
			response += fmt.Sprintf("%s", slot.Car.RegistrationNo)
		} else {
			response += fmt.Sprintf(", %s", slot.Car.RegistrationNo)
		}
	}
	return response, nil
}

func (p *ParkingService) FindAllocatedSlotByColour(colour string) (string, error) {
	slots, err := p.FindSlotByFeild(colour, "colour")
	if nil != err {
		return "", err
	}
	var response string
	for _, slot := range slots {
		if "" == response {
			response += fmt.Sprintf("%d", slot.Position)
		} else {
			response += fmt.Sprintf(", %d", slot.Position)
		}
	}
	return response, nil
}

func (p *ParkingService) FindSlotByRegistationNo(regNo string) (string, error) {
	slots, err := p.FindSlotByFeild(regNo, "registration_number")
	if nil != err {
		return "", err
	}
	var response string
	for _, slot := range slots {
		if "" == response {
			response += fmt.Sprintf("%d", slot.Position)
		} else {
			response += fmt.Sprintf(", %d", slot.Position)
		}
	}
	return response, nil
}
