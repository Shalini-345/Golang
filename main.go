package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 60
	var remainingTickets uint = 60
	var bookings []string
	/*
		fmt.Println("Welcome to our", conferenceName, "booking application")
		fmt.Println("we have the tickets total of ", conferenceTickets, "and we have ", remainingTickets, "tickets available")
		fmt.Println("get your tickets to attend")
	*/
	// formatting lines //

	fmt.Printf("Welcomr to our %v booking application \n", conferenceName)
	fmt.Printf("we have the tickets total of %v and we have %v tickets available\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets now to attend")

	for {
		var firstName string
		var lastName string
		var Email string
		var userTickets uint
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter your Email: ")
		fmt.Scan(&Email)

		fmt.Print("Enter no.of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v conformation email will be sent to %v shortly\n", firstName, lastName, userTickets, Email)
		fmt.Printf("we have %v left\n", remainingTickets)

		fmt.Printf(" these are all the bookings %v\n", bookings)
	}

}
