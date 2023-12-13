package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

// Package lever variables (accessible outside funkctions, e.g. "func main")
const conferenceTickets = 50         // How many tickets are available. "const" can not be changed!
var conferenceName = "Go Conference" // "var" is alowed to change.
var remainingTickets uint = 50       // uint declares that the variable is an int but cannot be negative!
// var bookings [conferenceTickets]string // Arrays: the square brackets [50] defines how many items the array can hold, here 50 items. Then we need to declare the type
// var bookings []string // Slice. It is like an array, but the size of it is dynamic! Use Slices rather than arrays if you don't know the max size of the list.
// var bookings = make([]map[string]string, 0) // Declares bookings as a Slice (the square brackets have no value, max size). We're declaring bookings as a list of maps. "map[string]string" means it is of the type map (like python dictionary), with the key and values as strings
// ", 0" at the end of make([]map[string]string, 0) says that the size of the list is 0, but will increase
var bookings = make([]UserData, 0) // creates an empty list of UserData structures

// We're creating a struct on this package level. Structs ("structures") can hold mixed data types, as oposed to maps
// "type" means that we are creating a custom type in our application, called "UserData", based on a structure of firstName, lastName, ...
// A "struct" can be compared to what other languages call "class", but its more light weight because "struct" does not support inheritance
type UserData struct {
	// properties of UserData:
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	// When assigning a var or const with a value simultaneously when declaring it, Go wil understand what type it is.
	// E.g. conferenceName = "Go Conference" is understood to be a string

	greetUsers()
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
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Booking logic:
			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)
			// Print first names of bookings, calling a function:
			firstNames := getFirstNames()
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend") // the ln in println indicates that a new line is added at the end of the print statement
}
func getFirstNames() []string { // Det som er inni () er input, etterpå output -> "[]string"
	firstNames := []string{}           // defines firstName as a "Slice" (dynamically sized array), indicated by the empty [], says its of type "string" and gives it an empty list {}
	for _, booking := range bookings { // replace index with _, it indicates that we have a variable, but its not used
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// OLD: create a map for a user
	// var userData = make(map[string]string) // "map" here, is a type -> map[data type of key]data type of the value
	// userData["firstName"] = firstName      // the first key/value pair we save to the userData map
	// userData["lastName"] = lastName        // the name within the [], the keys, can be named whatever you like
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) // in Go, one cannot mix key or value data types. This uint needs to be parsed as str. E.g. 30 tickets will be parsed as "30"

	// NEW: using struct in stead of map
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	// Kom til 2:17, må definere "Package Level Variables" (globale variabler)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	// for [deklarere og initialisere variabler i loopen];[betingelsen som bestemmer om loopen skal fortsette eller avslutte];[inkrement/dekrement, f.eks. øke variabel med 1]
	// time.Sleep(10 * time.Second)
	fmt.Printf("Sending confirmation email, please wait:\n")
	for sec := 9; sec > 0; sec-- {
		fmt.Printf("%v seconds left...\n", sec)
		time.Sleep(1 * time.Second)
	}
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //fmt.Sprintf returns a string instead of just printing the string
	fmt.Println("#################")
	fmt.Printf("Email with tickets successfully for:\n %v \nto email adress %v\n", ticket, email)
	fmt.Println("#################")
}
