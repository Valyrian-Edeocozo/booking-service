package main

import (
	"booking-app/src/helper"
	"fmt"
	"strconv"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets = 50
var bookings []string

func main() {

	var year = "1234"
	intValue, err := strconv.Atoi(year)

	if err != nil {
		panic("error")
	}

	value := intValue / 100

	result := strconv.Itoa(value + 1)

	var numValue = result[len(result)-1:]
	var suffix = ""

	switch numValue {
	case "1":
		suffix = "st"
	case "2":
		suffix = "nd"
	case "3":
		suffix = "rd"
	default:
		suffix = "th"

	}
	// return valueToReturn
	fmt.Println(result + suffix)

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
