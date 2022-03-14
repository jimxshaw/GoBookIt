package main

import (
	"fmt"
	"time"
)

func main() {

	eventName := "Go Conf"
	const totalTickets = 100
	remainingTickets := 100

	fmt.Println("GoBookIt is your premier booking app!")
	fmt.Printf("There are %v tickets total and %v tickets remaining\n", totalTickets, remainingTickets)
	fmt.Printf("Get tickets to %v %v here\n", eventName, time.Now().Year())

	var userName string
	var userTickets int

	// Ask user for name.
	userName = "James"
	userTickets = 4

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
