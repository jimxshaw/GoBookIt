package main

import (
	"strings"
)

func validateUserInput(userName, email string, userTickets uint) (bool, bool, bool) {
	isValidUserName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidUserName, isValidEmail, isValidTicketNumber
}
