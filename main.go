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

	// Two ways to print a text with variables:
	// fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("Welcome to %v booking application\n", conferenceName) // %v is a placeholder for the VALUE of the referenced variable
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend") // the ln in println indicates that a new line is added at the end of the print statement

	for {

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

		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}           // defines firstName as a "Slice" (dynamically sized array), indicated by the empty [], says its of type "string" and gives it an empty list {}
			for _, booking := range bookings { // replace index with _, it indicates that we have a variable, but its not used
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of the bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is book out. Come back next year.")
				break //breaks the indefinate for loop that keeps the application running
			}
		} else {

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}
}
