package Slot

import "FK_Machine_Coading/Workout"

type Slot struct {
	StartTime string
	EndTime string
	StartMeridiem string
	EndMeridiem  string
	WorkOuts []*Workout.WorkOut
}
