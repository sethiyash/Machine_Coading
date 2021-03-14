package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"FK_Machine_Coading/Center"
	"FK_Machine_Coading/Constants"
	"FK_Machine_Coading/FlipFit"
	"FK_Machine_Coading/User"
)

func main()  {

	reader := bufio.NewReader(os.Stdin)

	flipkartFitObj := new(FlipFit.FlipFit)

	flipkartFitObj.UserIdInit = 0
	flipkartFitObj.Users = make(map[string]*User.User)

	  flipkartFitObj.Centers = []*Center.Center{
		{
			Name: "Koramangala",
			Address: "Koramangala, Bengaluru",
			SlotsPerDay: Constants.KoramanglaSlots,
		},{
			Name: "Bellandur",
			SlotsPerDay: Constants.BellandurCenterSlots,
		},
	  }



	for {
		command ,_ := reader.ReadString('\n')
		command = strings.Trim(command, "\n")
		commands := strings.Split(command, " ")
		commandType := commands[0]

		switch commandType {
		case "register_user":
			if len(commands) ==  4 {
				flipkartFitObj.AddUser(commands[1], commands[2], commands[3])
			} else {
				fmt.Println("Invalid Input")
			}
		case "view_slots":
			if len(commands) == 2 {
				day, _ := strconv.Atoi(commands[1])
				flipkartFitObj.ViewAllSlots(day)
			}
		case "book_slot":
			day, _ := strconv.Atoi(commands[6])
			//userName, centerName, workOut, startTime, endTime string, day int
			flipkartFitObj.BookSlot(commands[1], commands[2], commands[3], commands[4], commands[5], day)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("default")

		}
	}
	return
}
