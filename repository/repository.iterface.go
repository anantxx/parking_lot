package repository

import "github.com/parking_lot/model"

type RepositoryInstance interface {
	GetParking() *model.Parking
	InitializeLot(int)
	/*	AllocateSlot(string, string) (string, error)
		ReleaseSlot(int) (string, error)
		ShowStatus() (string, error)
		FindRegistationNosByColour(string) (string, error)
		FindAllocatedSlotByColour(string) (string, error)
		FindSlotByRegistationNo(string) (string, error)*/
}
