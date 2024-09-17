package main

import (
	"fmt"
)

func greetings(conferenceName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
