package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	eventName := "Go Conf " + strconv.Itoa(time.Now().Year())
	const totalTickets = 100
	var remainingTickets uint = 100
	var bookings []string

	fmt.Println("GoBookIt is your premier booking app!")
	fmt.Printf("There are %v tickets total and %v tickets remaining\n", totalTickets, remainingTickets)
	fmt.Printf("Get tickets to %v here\n", eventName)

	for {
		var userName string
		var email string
		var userTickets uint

		// Get user input.
		fmt.Println("Enter your user name: ")
		fmt.Scan(&userName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets to book: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, userName)

		fmt.Printf("User %v booked %v tickets. Confirmation will be sent to the email %v.\n", userName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)

		fmt.Printf("Current bookings: %v\n", bookings)
	}

}
