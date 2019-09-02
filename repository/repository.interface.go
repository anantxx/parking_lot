package repository

import "github.com/parking_lot/model"

type RepositoryInstance interface {
	GetParking() *model.Parking
	InitializeLot(int)
	/*	//AllocateSlot(string, string) (string, error)
		ReleaseSlot(int) (string, error)
		ShowStatus() (string, error)
		FindRegistationNosByColour(string) (string, error)
		FindAllocatedSlotByColour(string) (string, error)
		FindSlotByRegistationNo(string) (string, error)*/
}

type SlotRepositoryInstance interface {
	AddNewSlot(*model.Slot, *model.Slot) int
	FindSlotBySlotNo(*model.Slot, int) (*model.Slot, error)
	ReleaseSlot(*model.Slot) error
	FindAllSlot(*model.Slot) []model.Slot
	FindSlotsByFeild(*model.Slot, string, string) []model.Slot
}
