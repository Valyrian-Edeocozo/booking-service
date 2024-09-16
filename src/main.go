package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	var bookings []string

	for {

		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets);
		fmt.Println("Get your tickets here to attend")
	
		var firstName string
		var lastName string
		var email string
		var userTickets int
	
		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
	
		fmt.Println("Enter your email address")
		fmt.Scan(&email)
	
		fmt.Println("Enter the number of tickets you want to buy")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
            fmt.Printf("Sorry, we only have %v tickets left\n", remainingTickets)
            continue
        }
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	
		remainingTickets = remainingTickets - userTickets
	
		fmt.Printf("The remaining ticket is %v\n", remainingTickets)

		bookings = append(bookings, firstName + " " + lastName)

		firstNames := []string{}

		for _, name := range bookings {
			var names = strings.Fields(name)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, all tickets are sold out. Come back next year")
            break
		}

	}
}