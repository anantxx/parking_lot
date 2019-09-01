package model

type Slot struct {
	PrevSlot *Slot
	Car      *Car
	Position int
	NextSlot *Slot
}
