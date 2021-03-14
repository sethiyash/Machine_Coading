package FlipFit

import (
	"fmt"

	"FK_Machine_Coading/Book"
	"FK_Machine_Coading/Center"
	"FK_Machine_Coading/User"
)

type FlipFit struct {
	Users map[string]*User.User
	Centers []*Center.Center
	UserIdInit int
}

func (f *FlipFit) AddUser(name, email, contact string) {
	if f.Users[name] != nil {
		user := &User.User{
			Id:            f.UserIdInit + 1,
			Name:          name,
			Email:         email,
			ContactNumber: contact,
			BookedSlots:   nil,
		}
		f.Users[name] = user
		fmt.Println("Thank You! You are now member of FlipFit")
	} else {
		fmt.Println("Y")
	}

}

func (f *FlipFit) ViewAllSlots(day int) {
	for _, center := range f.Centers {
		slots := center.SlotsPerDay[day]
		if len(slots) > 0 {
			for _, slot := range slots {
				for _, workout := range slot.WorkOuts {
					if workout.NoOfSeats > 0 {
						fmt.Printf("WorkOut - %s is available at time %s - %s \n", workout.Name,
							slot.StartTime, slot.EndTime)
					}
				}
			}
		}
	}
}

func (f *FlipFit) BookSlot(userName, centerName, workOut, startTime, endTime string, day int) {
	if f.Users[userName] != nil {
		isCenterAvailable := false
		for _, center := range f.Centers {
			if center.Name == centerName {
				isCenterAvailable = true
				for _, slots := range center.SlotsPerDay[day] {
					if slots.StartTime == startTime && slots.EndTime == endTime {
						for _, workout := range slots.WorkOuts {
							if workout.Name == workOut {
								if workout.NoOfSeats > 0 {
									workout.NoOfSeats--
									f.Users[userName].BookedSlots = []*Book.Book{
										{
											CenterName: centerName,
											StartTime:  startTime,
											EndTime:    endTime,
											Day:        day,
										},
									}
								} else {
									fmt.Println("No seats available")
								}
							}
						}
					}
				}
			}
		}
		if !isCenterAvailable {
			fmt.Println("There is no given center available")
		}
	} else {
		fmt.Println("User does not exists")
	}
}