package main

import (
	"booking-app/src/helper"
	"fmt"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets = 50
var bookings []string

func main() {

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets int

		greetings(conferenceName, conferenceTickets, remainingTickets)
		entries.CollectInput(firstName, lastName, email, userTickets)
		entries.ValidateEntries(firstName, lastName, email, userTickets, remainingTickets)
		firstNames := entries.FormatNames(bookings, firstName, lastName)

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	}
}
