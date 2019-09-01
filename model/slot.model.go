package model

type Slot struct {
	prevSlot *Slot
	car      *Car
	position int
	nextSlot *Slot
}
