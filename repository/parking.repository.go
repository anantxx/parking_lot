package repository

import (
	"sync"

	"github.com/parking_lot/model"
)

var instance *model.Parking
var once sync.Once

type ParkingRepository struct {
	Parking *model.Parking
}

func GetInstance() *model.Parking {
	once.Do(func() {
		instance = &model.Parking{}
	})
	return instance
}

func NewParkingRepository(parking *model.Parking) ParkingRepository {
	return ParkingRepository{Parking: parking}
}

func (p ParkingRepository) GetParking() *model.Parking {
	return p.Parking
}


func (p ParkingRepository) InitializeLot(totalSlot int)(error) {
	p.Parking.TotalSlot = totalSlot
	p.Parking.IsParkingLotCreated = true
	return nil
}
