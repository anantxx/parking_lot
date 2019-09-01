package model

type Parking struct {
	TotalSlot           int
	IsParkingLotCreated bool
	AllocatedSlot       int
	Slots               *Slot
}
