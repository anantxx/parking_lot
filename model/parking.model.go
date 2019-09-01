package model

type Parking struct {
	totalSlot           int
	isParkingLotCreated bool
	allocatedSlot       int
	slots               *Slot
}
