package service

import (
	"fmt"
	"strings"

	. "github.com/parking_lot/constant"
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

	if p.ParkingRepository.GetParking().TotalSlot == p.ParkingRepository.GetParking().AllocatedSlot {
		return "", ERR_PARKING_FULL
	}

	car := repository.NewCar(regNo, colour)
	var pos int
	if 0 == p.ParkingRepository.GetParking().AllocatedSlot {
		fmt.Println("INISDE FIRST")
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
