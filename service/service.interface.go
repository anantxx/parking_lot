package service

import "github.com/parking_lot/model"

//ServicesInstance Inteface
type ServicesInstance interface {
	InitializeLot(int) (string, error)
	AllocateSlot(string, string) (string, error)
	ReleaseSlot(int) (string, error)
	ShowStatus() (string, error)
	FindRegistationNosByColour(string) (string, error)
	FindAllocatedSlotByColour(string) (string, error)
	FindSlotByRegistationNo(string) (string, error)
	FindSlotByFeild(string, string) ([]model.Slot, error)
}
