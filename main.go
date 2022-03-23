package main

import (
	"fmt"
	"gobookit/helpers"
	"strconv"
	"sync"
	"time"
)

const totalTickets = 100

var eventName string = "Go Conf " + strconv.Itoa(time.Now().Year())
var remainingTickets uint = 100

// Create initial empty Slice of Maps.
var bookings = make([]UserData, 0)

type UserData struct {
	userName    string
	email       string
	userTickets uint
	eventName   string
}

// Waits for launched goroutines to finish.
var waitGroup = sync.WaitGroup{}

func main() {

	greetUsers()

	userName, email, userTickets := getUserInput()

	isValidUserName, isValidEmail, isValidTicketNumber := helpers.ValidateUserInput(userName, email, userTickets, remainingTickets)

	if isValidUserName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, eventName, userName, email)

		// Sets the number of goroutines to wait for.
		waitGroup.Add(1)

		// Make this function call concurrent.
		// The number of goroutines is the number to
		// be added to the wait group.
		go sendTicket(userTickets, eventName, userName, email)

		// Get usernames.
		userNames := getUserNames()
		fmt.Printf("Usernames of current bookings: %v\n", userNames)

		atBookingCapacity := remainingTickets == 0

		if atBookingCapacity {
			// End the program.
			fmt.Printf("%v is at capacity bookings. Please return next year!\n", eventName)
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

	// Wait should be called at the end
	// of the main thread. This blocks the
	// main thread until wait group counter is 0.
	// Essentially main thread waits until all
	// the theads that have been added with Add
	// are done before the Application can exit.
	waitGroup.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", eventName)
	fmt.Printf("There are %v tickets total and %v tickets remaining\n", totalTickets, remainingTickets)
	fmt.Println("Get tickets here!")
}

func getUserNames() []string {
	userNames := []string{}

	for _, booking := range bookings {
		userNames = append(userNames, booking.userName)
	}

	return userNames
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

	// Create a Map for a user.
	var userData = UserData{
		userName:    userName,
		email:       email,
		eventName:   eventName,
		userTickets: userTickets,
	}

	// Add Map to Slice.
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("User %v booked %v tickets. Confirmation will be sent to the email %v.\n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, eventName)
}

func sendTicket(userTickets uint, eventName, userName, email string) {
	// Simulate delay in sending ticket.
	time.Sleep(5 * time.Second)

	var ticket = fmt.Sprintf("%v %v tickets for %v\n", userTickets, eventName, userName)
	fmt.Println("************************************")
	fmt.Printf("Sending ticket:\n %v to email %v\n", ticket, email)
	fmt.Println("************************************")

	// Done decrements the wait group counter by 1.
	// This is called by the goroutine to show
	// that it's finished.
	// Once this is finished then the main thread
	// wouldn't have to wait anymore.
	waitGroup.Done()
}
