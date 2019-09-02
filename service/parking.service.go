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

	p.ParkingRepository.InitializeLot(totalSlot)
	response := fmt.Sprintf("Created a parking lot with %d slots", p.ParkingRepository.GetParking().TotalSlot)
	return response, nil
}

func (p *ParkingService) AllocateSlot(regNo string, colour string) (string, error) {

	if false == p.ParkingRepository.GetParking().IsParkingLotCreated {
		return "", ERR_NO_INITIALIZATION
	}

	if "" == strings.Trim(regNo, " ") || "" == strings.Trim(colour, " ") {
		return "", ERR_INVALID_ARGUMENT
	}

	if p.ParkingRepository.GetParking().TotalSlot <= p.ParkingRepository.GetParking().AllocatedSlot {
		return "", ERR_PARKING_FULL
	}

	car := repository.NewCar(regNo, colour)
	var pos int
	if 0 == p.ParkingRepository.GetParking().AllocatedSlot {
		slot := repository.NewSlot(car, 1)
		p.ParkingRepository.GetParking().Slots = slot
		pos = slot.Position
	} else {
		newSlot := repository.NewSlot(car, 0)
		pos = p.SlotRepository.AddNewSlot(p.ParkingRepository.GetParking().Slots, newSlot)
	}
	p.ParkingRepository.GetParking().AllocatedSlot++
	return fmt.Sprintf("Allocated slot number: %d", pos), nil
}

func (p *ParkingService) ReleaseSlot(slotNo int) (string, error) {
	if false == p.ParkingRepository.GetParking().IsParkingLotCreated {
		return "", ERR_NO_INITIALIZATION
	}

	if p.ParkingRepository.GetParking().AllocatedSlot < 1 {
		return "", ERR_NO_CAR_PARKED
	}

	realeaseSlot, err := p.SlotRepository.FindSlotBySlotNo(p.ParkingRepository.GetParking().Slots, slotNo)
	if nil != err {
		return "", err
	}

	err = p.SlotRepository.ReleaseSlot(realeaseSlot)
	if nil != err {
		return "", err
	}
	p.ParkingRepository.GetParking().AllocatedSlot--
	return fmt.Sprintf("Slot number %d is free", realeaseSlot.Position), nil

}

func (p *ParkingService) ShowStatus() (string, error) {
	if false == p.ParkingRepository.GetParking().IsParkingLotCreated {
		return "", ERR_NO_INITIALIZATION
	}

	if p.ParkingRepository.GetParking().AllocatedSlot < 1 {
		return "", ERR_NO_CAR_PARKED
	}

	slots := p.SlotRepository.FindAllSlot(p.ParkingRepository.GetParking().Slots)
	response := fmt.Sprintf("Slot No.\tRegistration No\tColor")
	for _, slot := range slots {
		response += fmt.Sprintf("\n%d\t\t%s\t%s", slot.Position, slot.Car.RegistrationNo, slot.Car.Colour)
	}
	return response, nil
}

func (p *ParkingService) FindSlotByFeild(colour string, feild string) ([]model.Slot, error) {
	if false == p.ParkingRepository.GetParking().IsParkingLotCreated {
		return []model.Slot{}, ERR_NO_INITIALIZATION
	}

	if p.ParkingRepository.GetParking().AllocatedSlot < 1 {
		return []model.Slot{}, ERR_NO_CAR_PARKED
	}

	slots := p.SlotRepository.FindSlotsByFeild(p.ParkingRepository.GetParking().Slots, colour, feild)
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
