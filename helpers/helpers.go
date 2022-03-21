package helpers

import (
	"strings"
)

func ValidateUserInput(userName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidUserName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidUserName, isValidEmail, isValidTicketNumber
}
