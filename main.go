package main

import "fmt"

func main() {

	// When assigning a var or const with a value simultaneously when declaring it, Go wil understand what type it is.
	// E.g. conferenceName = "Go Conference" is understood to be a string
	var conferenceName = "Go Conference" // "var" is alowed to change.
	const conferenceTickets = 50         // How many tickets are available. "const" can not be changed!
	var remainingTickets uint = 50       // uint declares that the variable is an int but cannot be negative!

	// Two ways to print a text with variables:
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName) // %v is a placeholder for the VALUE of the referenced variable

	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend") // the ln in println indicates that a new line is added at the end of the print statement

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
