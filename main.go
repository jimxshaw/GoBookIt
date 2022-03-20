package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const totalTickets = 100

var eventName string = "Go Conf " + strconv.Itoa(time.Now().Year())
var remainingTickets uint = 100
var bookings []string

func main() {

	greetUsers()

	for {
		userName, email, userTickets := getUserInput()

		isValidUserName, isValidEmail, isValidTicketNumber := validateUserInput(userName, email, userTickets)

		if isValidUserName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, eventName, userName, email)

			// Get usernames.
			userNames := getUserNames()
			fmt.Printf("Usernames of current bookings: %v\n", userNames)

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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", eventName)
	fmt.Printf("There are %v tickets total and %v tickets remaining\n", totalTickets, remainingTickets)
	fmt.Println("Get tickets here!")
}

func getUserNames() []string {
	userNames := []string{}

	for _, booking := range bookings {
		var userInfo = strings.Fields(booking)
		userNames = append(userNames, userInfo[0])
	}

	return userNames
}

func validateUserInput(userName, email string, userTickets uint) (bool, bool, bool) {
	isValidUserName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidUserName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, uint) {
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

	return userName, email, userTickets
}

func bookTicket(userTickets uint, eventName, userName, email string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, userName+" "+email)

	fmt.Printf("User %v booked %v tickets. Confirmation will be sent to the email %v.\n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}
