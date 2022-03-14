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
	var email string
	var userTickets int

	// Get user input.
	fmt.Println("Enter your user name: ")
	fmt.Scan(&userName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to book: ")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets. Confirmation will be sent to the email %v.\n", userName, userTickets, email)

}
