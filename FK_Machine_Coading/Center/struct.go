package Center

import "FK_Machine_Coading/Slot"

type Center struct {
	Name string
	Address string
	SlotsPerDay map[int][]*Slot.Slot
}