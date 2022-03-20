package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	eventName := "Go Conf " + strconv.Itoa(time.Now().Year())
	const totalTickets = 100
	var remainingTickets uint = 100
	var bookings []string

	greetUsers(eventName, totalTickets, remainingTickets)

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

		isValidUserName := len(userName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidUserName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, userName+" "+email)

			fmt.Printf("User %v booked %v tickets. Confirmation will be sent to the email %v.\n", userName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)

			// Print usernames.
			printUserNames(bookings)

			atBookingCapacity := remainingTickets == 0

			if atBookingCapacity {
				// End the program.
				fmt.Printf("%v is at capacity bookings. Please return next year!\n", eventName)
				break
			}
		} else {
			if !isValidEmail {
				fmt.Println("Invalid email. Try again.")
			}

			if !isValidUserName {
				fmt.Println("Invalid username. Try again.")
			}

			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number. Try again.")
			}

		}

	}

}

func greetUsers(eventName string, totalTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", eventName)
	fmt.Printf("There are %v tickets total and %v tickets remaining\n", totalTickets, remainingTickets)
	fmt.Println("Get tickets here!")
}

func printUserNames(bookings []string) {
	userNames := []string{}

	for _, booking := range bookings {
		var userInfo = strings.Fields(booking)
		userNames = append(userNames, userInfo[0])
	}

	fmt.Printf("Usernames of current bookings: %v\n", userNames)
}
