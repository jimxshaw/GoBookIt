package main

import "fmt"

func main() {

	var eventName = "Go Conf"
	const eventTickets = 100
	var remainingTickets = 100

	fmt.Println("GoBookIt is your premier booking app!")
	fmt.Println("There are", eventTickets, "tickets total and", remainingTickets, "tickets remaining")
	fmt.Println("Get tickets to", eventName, "here")

}
