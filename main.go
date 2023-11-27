package main

import (
	"fmt"
	"strings"
)

func main() {

	// When assigning a var or const with a value simultaneously when declaring it, Go wil understand what type it is.
	// E.g. conferenceName = "Go Conference" is understood to be a string
	var conferenceName = "Go Conference" // "var" is alowed to change.
	const conferenceTickets = 50         // How many tickets are available. "const" can not be changed!
	var remainingTickets uint = 50       // uint declares that the variable is an int but cannot be negative!
	// var bookings [conferenceTickets]string // Arrays: the square brackets [50] defines how many items the array can hold, here 50 items. Then we need to declare the type
	// var bookings []string // Slice. It is like an array, but the size of it is dynamic! Use Slices rather than arrays if you don't know the max size of the list.
	bookings := []string{} // Declares bookings as a Slice (the square brackets have no value, max size), and the "array" of the slice is empty "{}"

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	// Two ways to print a text with variables:
	// fmt.Println("Welcome to", conferenceName, "booking application")

	for {

		// Example of OR statement:
		// isValidCity := city == "singapore" || city == "London"
		//
		// Also, we can tur the "postive check" around, by using ! in front of the variable, like this:
		// !isValidCity
		//
		// That would equal to:
		// isValidCity := city != "singapore" || city != "London"
		//
		// So now we could use !isValidCity to check if the "city" variable is NOT LIKE "singapore" OR "London"
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Booking logic:
			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			// Print first names of bookings, calling a function:
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is book out. Come back next year.")
				break //breaks the indefinate for loop that keeps the application running
			}

		} else {
			if !isValidName { // translates to IF isINVALIDname
				fmt.Printf("The name you entered seems wrong. You typed %v %v\n", firstName, lastName)
			}
			if !isValidEmail { // translates to IF isINVALIDname
				fmt.Printf("The email you entered is invalid. You typed %v\n", email)
			}
			if !isValidTicketNumber { // translates to IF isINVALIDname
				fmt.Printf("The number of tickets you entered is invalid. You typed %v. Available tickets: %v.\n", userTickets, remainingTickets)
			}
			// continue is not nesescary when it is the bottom of a for loop
		}
	}

}

func greetUsers(confName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend") // the ln in println indicates that a new line is added at the end of the print statement
}
func getFirstNames(bookings []string) []string { // Det som er inni () er input, etterpå output -> "[]string"
	firstNames := []string{}           // defines firstName as a "Slice" (dynamically sized array), indicated by the empty [], says its of type "string" and gives it an empty list {}
	for _, booking := range bookings { // replace index with _, it indicates that we have a variable, but its not used
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	// This is actually a short for:
	// var isValidName bool = len(firstName) >= 2 && len(lastName) >=2
	// because Go understands its a declaration of a variable of the bool type:
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets // AND = &&
	return isValidName, isValidEmail, isValidTicketNumber
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) // the Scan method will ask for user input and store it in the referenced variable. But we need to add a "pointer" to that variable (the % sign).
	// A "pointer" is  pointing to the RAM adress of the referenced variable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your first email adress: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// Kom til 2:17, må definere "Package Level Variables" (globale variabler)
}
