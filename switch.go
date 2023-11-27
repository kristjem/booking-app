package main

import (
	"fmt"
)

func switch_test() {
	city := "London"

	switch city {
	case "New Your":
		// execute code for booking New York conference tickets
	case "Singapore", "Hong Kong":
		// execute code for booking Singapore conference tickets
	case "London", "Berlin":
		// execute code for booking London and Berlin conference tickets
		fmt.Printf("City is: %v", city)
	case "Mexico City":

	default:
		fmt.Println("No valid city selected")
	}
}

func main() {
	switch_test()
}
