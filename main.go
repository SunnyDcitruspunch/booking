package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to come in")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int
		bookings := []string{}

		// ask user for their name
		// fmt.Println(&remainingTickets)
		// a pointer!
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you want: ")
		fmt.Scan(&userTickets)

		if userTickets > int(remainingTickets) {
			fmt.Printf("We only have %[1]v tickets left, please book no more than %[1]v tickets.\n", remainingTickets)
			break
		}

		// Add the full name to the bookings slice
		fullName := firstName + " " + lastName
		bookings = append(bookings, fullName)

		remainingTickets = conferenceTickets - uint(userTickets)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Booked user names: %v\n", firstNames)
		// fmt.Printf("Booked first name: %v\n", bookings[0])
		// fmt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Array length: %v \n", len(bookings))

		var noRemainingTickets bool = remainingTickets == 0
		// alternative syntax: noRemainingTickets := remainingTickets == 0
		if noRemainingTickets {
			// end program
			fmt.Println("Our conference is booked out. Come back next year :)")
			break
		}
	}
}
