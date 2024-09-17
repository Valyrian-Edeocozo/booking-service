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
		firstName, lastName, email, userTickets = entries.CollectInput(firstName, lastName, email, userTickets)
		var ticketsRemaining = &remainingTickets
		entries.ValidateEntries(firstName, lastName, email, userTickets, ticketsRemaining)

		bookings = append(bookings, firstName+" "+lastName)

		firstNames := entries.FormatNames(bookings, firstName, lastName)

		fmt.Printf("The first names of bookings are: %v\n", firstNames)

	}
}
