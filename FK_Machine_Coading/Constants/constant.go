package Constants

import (
	"FK_Machine_Coading/Slot"
	"FK_Machine_Coading/Workout"
)

var (
	WorksOutsData = []*Workout.WorkOut{
		{
			Name: "Yoga",
			Details: "",
			NoOfSeats: 2,
		},
		{
			Name: "Weights",
			NoOfSeats: 2,
		},
		{
			Name: "Cardio",
			NoOfSeats: 2,
		},
	}


	KoramanglaSlots = map[int][]*Slot.Slot{
		1: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},

		},
		2: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},

		},
		3 : []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},

		},
		4 : []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},

		},
		5 : []*Slot.Slot{
				{
					StartTime: "6AM",
					EndTime: "7AM",
					WorkOuts: WorksOutsData,
				},{
					StartTime: "7AM",
					EndTime: "8AM",
					WorkOuts: WorksOutsData,
				},{
					StartTime: "8AM",
					EndTime: "9AM",
					WorkOuts: WorksOutsData,
				},
		},
		6 : []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},

		},
	}

	BellandurCenterSlots = map[int][]*Slot.Slot{
		1: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},
		},
		2: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},
		},
		3: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},
		},
		4: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},
		},
		5: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},
		},
		6: []*Slot.Slot{
			{
				StartTime: "6AM",
				EndTime: "7AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "7AM",
				EndTime: "8AM",
				WorkOuts: WorksOutsData,
			},{
				StartTime: "8AM",
				EndTime: "9AM",
				WorkOuts: WorksOutsData,
			},
		},
	}


)
