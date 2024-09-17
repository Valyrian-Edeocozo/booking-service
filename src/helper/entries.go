package entries

import (
	"fmt"
	"strings"
)

func CollectInput(firstName string, lastName string, email string, userTickets int) (string, string, string, int) {
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to buy")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func ValidateEntries(firstName string, lastName string, email string, userTickets int, remainingTickets *int) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmailAddress := strings.Contains(email, "@")
	isValidNumberOfTickets := userTickets <= *remainingTickets && userTickets > 0

	if isValidEmailAddress && isValidNumberOfTickets && isvalidName {

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		*remainingTickets = *remainingTickets - userTickets

		fmt.Printf("The remaining ticket is %v\n", *remainingTickets)

		if *remainingTickets == 0 {
			fmt.Println("Sorry, all tickets are sold out. Come back next year")
		}
	} else {
		if !isValidEmailAddress {
			fmt.Println("Invalid email address. Please enter a valid email address.")
		}
		if !isValidNumberOfTickets {
			fmt.Println("Invalid number of tickets")
		}
		if !isvalidName {
			fmt.Println("Invalid name. First and last names should be at least 2 characters long.")
		}
	}
}
