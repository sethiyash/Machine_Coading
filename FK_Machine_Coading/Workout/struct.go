package Workout

import "FK_Machine_Coading/User"

type WorkOut struct {
	Name string
	Details string
	NoOfSeats int
	Users []*User.User
}
