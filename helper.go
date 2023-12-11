package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// This is actually a short for:
	// var isValidName bool = len(firstName) >= 2 && len(lastName) >=2
	// because Go understands its a declaration of a variable of the bool type:
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // AND = &&
	return isValidName, isValidEmail, isValidTicketNumber
}
